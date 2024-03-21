from dotenv import load_dotenv
import geopandas as gpd
from upsert.conn import DatabaseAccess

from geoalchemy2 import Geometry, WKTElement
from sqlalchemy import Integer, String, Float


from typing import Final


if __name__ == "__main__":
    load_dotenv(".env")

    PNU: Final = "~/Developer/gis-app/data/asset/LSMD_CONT_LDREG_SEOUL/LSMD_CONT_LDREG_11_202402.shp"

    db = DatabaseAccess()
    data = gpd.read_file(PNU, encoding='cp949').to_crs(epsg=4326)
    data.columns = ['id', 'jibun', 'bchk', "pnu", "col_adm_se", "geometry"]
    column = {
        "SGG_OID": String,
        "JIBUN": String,
        "BCHK": String,
        "PNU": String,
        "COL_ADM_SE": String,
        "geometry": Geometry("GEOMETRY", srid=4326)
    }

    db.insert_geo_dataframe(
        data,
        "__TABLE_PNU",
        column,
        "replace",
    )
