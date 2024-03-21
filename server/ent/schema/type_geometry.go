package schema

import (
	"database/sql/driver"
	"fmt"
	"log"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"github.com/paulmach/orb"
	"github.com/paulmach/orb/encoding/wkb"
)

type Geometry struct {
	orb.Geometry
}

// Scan implements the Scanner interface for database/sql.
func (g *Geometry) Scan(value any) error {
	bin, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("invalid binary value for geometry: expected []bytes, got %T", value)
	}
	geom, err := wkb.Unmarshal(bin)
	if err != nil {
		log.Printf("Raw WKB data: %x, %s", bin, string(bin))
		return fmt.Errorf("geometry scan error: unable to unmarshal WKB, %v", err)
	}
	*g = Geometry{geom}
	return nil
}

// Value implements the driver Valuer interface.
func (g Geometry) Value() (driver.Value, error) {
	return wkb.Value(g.Geometry).Value()
}

// FormatParam implements the sql.ParamFormatter interface to format the SQL placeholder.
func (g Geometry) FormatParam(placeholder string, info *sql.StmtInfo) string {
	switch info.Dialect {
	case dialect.Postgres:
		return "ST_GeomFromEWKB(" + placeholder + ")"
	case dialect.MySQL:
		return "ST_GeomFromWKB(" + placeholder + ")"
	}
	return placeholder
}

// SchemaType defines the schema-type of the Geometry object for the database.
func (Geometry) SchemaType() map[string]string {
	return map[string]string{
		dialect.MySQL:    "GEOMETRY", // Generic geometry type
		dialect.Postgres: "geometry", // PostGIS also uses a generic type
	}
}
