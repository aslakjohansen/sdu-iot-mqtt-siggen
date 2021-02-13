produce:
	go run mqtt-siggen.go

subscribe:
	mosquitto_sub -v -t "siggen/+"

