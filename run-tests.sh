#!/bin/bash

if [[ -z "${DIREKTIV_TEST_URL}" ]]; then
	echo "Test URL is not set, setting it to http://localhost:9191"
	DIREKTIV_TEST_URL="http://localhost:9191"
fi

if [[ -z "${DIREKTIV_SECRET_sshcert}" ]]; then
	echo "Secret sshcert is required, set it with DIREKTIV_SECRET_sshcert"
	exit 1
fi

if [[ -z "${DIREKTIV_SECRET_sshpwd}" ]]; then
	echo "Secret sshpwd is required, set it with DIREKTIV_SECRET_sshpwd"
	exit 1
fi

docker run --network=host -v `pwd`/tests/:/tests direktiv/karate java -DtestURL=${DIREKTIV_TEST_URL} -Dlogback.configurationFile=/logging.xml -Dsshcert="${DIREKTIV_SECRET_sshcert}" -Dsshpwd="${DIREKTIV_SECRET_sshpwd}"  -jar /karate.jar /tests/v1.0/karate.yaml.test.feature ${*:1}