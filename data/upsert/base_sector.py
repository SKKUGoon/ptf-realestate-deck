import geopandas as gpd

from typing import Final


def upload(file: str):
    data = gpd.read_file(file)

    return data


if __name__ == "__main__":
    ASSET: Final = "~/Developer/gis-app/data/asset/Z_KAIS_TL_KODIS_BAS_11000/Z_KAIS_TL_KODIS_BAS_11000.shp"
    d = upload(ASSET)
    print(d)
    print(len(d))

    PNU: Final = "~/Developer/gis-app/data/asset/LSMD_CONT_LDREG_SEOUL/LSMD_CONT_LDREG_11_202402.shp"
    d2 = upload(PNU)
    print(d2)
    print(len(d2))
    ...
