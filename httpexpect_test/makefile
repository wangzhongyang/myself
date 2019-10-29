do_test:
	sh ./integration_testing.sh &
	sh ./unit_testing.sh &

do_integration_test:
	sh ./integration_testing.sh &

do_unit_test:
	sh ./unit_testing.sh &

stop_test:
	sh ./stop_test.sh 1
	sh ./stop_test.sh 2
stop_integration_test:
	sh ./stop_test.sh 1

stop_unit_test:
	sh ./stop_test.sh 2

clear_db:
	sh clear_db.sh province_warehouse_test