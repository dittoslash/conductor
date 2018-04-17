pipeline {
  agent any
  stages {
    stage('Build') {
      parallel {
        stage('Compile ARM') {
          steps {
            node(label: 'rpi') {
              git(url: 'https://github.com/dittoslash/conductor.git', branch: 'master')
              sh '''cd client
/usr/local/go/bin/go build'''
              archiveArtifacts 'conductor'
            }

          }
        }
        stage('Compile X86') {
          steps {
            sh '''cd client
/usr/local/go/bin/go build'''
            archiveArtifacts 'conductor'
          }
        }
      }
    }
  }
}