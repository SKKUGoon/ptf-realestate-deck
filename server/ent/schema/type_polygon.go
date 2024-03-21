package schema

import (
	"fmt"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"github.com/paulmach/orb"
	"github.com/paulmach/orb/encoding/wkb"
)

type Polygon struct {
	orb.Polygon
}

func (p *Polygon) Scan(value any) error {
	bin, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("invalid binary value for polygon")
	}

	var op orb.Polygon

	if err := wkb.Scanner(&op).Scan(bin); err != nil {
		return err
	}
	*p = Polygon{op}
	return nil
}

func (p Polygon) FormatParam(placeholder string, info *sql.StmtInfo) string {
	switch info.Dialect {
	case dialect.Postgres:
		return "ST_GeomFromEWKB(" + placeholder + ")"
	case dialect.MySQL:
		return "ST_GeomFromWKB(" + placeholder + ")"
	}
	return placeholder
}

func (Polygon) SchemaType() map[string]string {
	return map[string]string{
		dialect.MySQL:    "POLYGON",
		dialect.Postgres: "geometry(Polygon,4326)",
	}
}
