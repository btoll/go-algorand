pipeline {
    environment {
        AWS_ACCESS_KEY_ID = credentials("AWS_ACCESS_KEY_ID")
        AWS_SECRET_ACCESS_KEY = credentials("AWS_SECRET_ACCESS_KEY")
    }

    agent any

    stages {
        stage("test") {
            steps {
                sh 'NETWORK=mainnet S3_SOURCE=algorand-internal/channel VERSION=2.4.1 mule -f package-sign.yaml package-sign'
                sh 'BRANCH=rel/stable NETWORK=mainnet SHA=573a34c457e9ac5e4e4f63355f0b295ced001083 VERSION=2.4.1 mule -f package-test.yaml package-test'
            }
        }
    }
}

