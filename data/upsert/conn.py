import pandas as pd
import geopandas as gpd
from sqlalchemy import create_engine, MetaData, Table
from sqlalchemy.dialects.postgresql import insert
from geoalchemy2 import WKTElement
import os
from typing import Literal, List, Dict


class DatabaseAccess:
    def __init__(self, verbose: bool = False):
        username = os.getenv('__USERNAME')
        password = os.getenv('__PASSWORD')
        host = os.getenv('__HOSTNAME')
        port = os.getenv('__PORTNO')
        name = os.getenv('__DBNAME')

        # Create an engine
        self.engine_str = f'postgresql://{username}:{password}@{host}:{port}/{name}'
        if verbose:
            print(f"Connecting to {self.engine_str}")

    @staticmethod
    def _retrieve_table_name(value: str):
        table = os.getenv(value)
        if table == "" or table is None:
            raise KeyError(f"no {value} was found in env")

        if '.' in table:
            # It contains a schema
            return table.split('.')
        else:
            return None, table

    def insert_row(self,
                   records: List[Dict],
                   on_conflict_row: List[str],
                   target_table: str,
                   target_schema: str):
        # Insert on postgres database
        # If conflict, update
        engine = create_engine(self.engine_str)
        metadata = MetaData()

        table = Table(target_table, metadata, autoload_with=engine, schema=target_schema)

        inserted_ids = []
        with engine.connect() as conn:
            for record in records:
                stmt = insert(table).values(record)

                if len(on_conflict_row) > 0:
                    do_update_stmt = stmt.on_conflict_do_update(
                        index_elements=on_conflict_row,
                        set_=record
                    ).returning(table.c.id)  # Add returning clause to capture the id
                    result = conn.execute(do_update_stmt)
                    id_of_inserted_row = result.scalar()
                    inserted_ids.append(id_of_inserted_row)
                    # print(record, "executed", do_update_stmt)
                else:
                    stmt = stmt.returning(table.c.id)  # Add returning clause to capture the id
                    result = conn.execute(stmt)
                    id_of_inserted_row = result.scalar()
                    inserted_ids.append(id_of_inserted_row)
                    # print(record, "executed", stmt)
                conn.commit()
        engine.dispose()
        return inserted_ids

    def insert_dataframe(self,
                         data: pd.DataFrame,
                         target_table: str,
                         column_type: dict,
                         if_exists: Literal['append', 'replace'],
                         verbose: bool = True):
        schema, table = self._retrieve_table_name(target_table)

        engine = create_engine(self.engine_str)
        if column_type is None:
            data.to_sql(
                table,
                engine,
                schema=schema,
                if_exists=if_exists,
                index=False
            )
        else:
            data.to_sql(
                table,
                engine,
                schema=schema,
                if_exists=if_exists,
                index=False,
                dtype=column_type,
            )
        if verbose:
            print(f"Inserted {len(data)} rows into {table}")
        engine.dispose()

    def insert_geo_dataframe(self,
                             data: gpd.GeoDataFrame,
                             target_table: str,
                             column_type: dict,
                             if_exists: Literal['append', 'replace'],
                             geometry_column: str = 'geometry',
                             epsg: int = 4326,
                             verbose: bool = True):
        # Data column type check
        assert geometry_column in column_type.keys()
        schema, table = self._retrieve_table_name(target_table)

        # Before using `to_sql`, convert the geometry column in your DataFrame to WKT(Well-Known Text)
        # using element (WKTElement). This is necessary because geoalchemy2 expects geometries in WKT format
        # Should be single geometry column.
        df = data.copy(deep=True)
        df[geometry_column] = df[geometry_column].apply(
            lambda x: WKTElement(x.wkt, srid=epsg)
        )

        engine = create_engine(self.engine_str)
        df.to_sql(table, con=engine, schema=schema, if_exists=if_exists, index=False, dtype=column_type)
        if verbose:
            print(f"Inserted {len(data)} rows into {table}")
        engine.dispose()

    def select_dataframe(self, target_table: str, verbose: bool = True):
        schema, table = self._retrieve_table_name(target_table)

        query = f"SELECT * FROM {schema}.{table}"
        engine = create_engine(self.engine_str)

        if verbose:
            print(f"Sending query {query}")
        df = pd.read_sql(query, engine)
        engine.dispose()
        return df

    def select_geo_dataframe(self, target_table: str, geometry_column: str = 'geometry', verbose: bool = True):
        schema, table = self._retrieve_table_name(target_table)

        query = f"SELECT * FROM {schema}.{table}"
        engine = create_engine(self.engine_str)

        if verbose:
            print(f"Sending geo query {query}")
        gdf = gpd.read_postgis(query, engine, geom_col=geometry_column)
        engine.dispose()
        return gdf

