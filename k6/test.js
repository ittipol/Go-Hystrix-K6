import http from 'k6/http'

export let options = {
    stages:[
        {target: 3, duration: '5s'},
        {target: 4, duration: '4s'},
        {target: 5, duration: '8s'},
    ]
}

export default function() {
    http.get("http://host.docker.internal:8080/user")
}