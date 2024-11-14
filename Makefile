

build_store_image:
	docker build -t "store_api:0.0.1" -f ./images/Store.dockerfile .


build_warehouse_image:
	docker build -t warehouse_api:0.0.1 -f ./images/Warehouse.dockerfile .