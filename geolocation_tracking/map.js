let map = L.map('map').setView([1.469380, 124.844330], 12)

L.tileLayer('https://tile.openstreetmap.org/{z}/{x}/{y}.png', {
    maxZoom: 19,
    attribution: '&copy; <a href="http://www.openstreetmap.org/copyright">OpenStreetMap</a>'
}).addTo(map);

let mikroIcon = L.icon({
    iconUrl: "mikrolet.png",
    iconSize: [50, 50],
    iconAnchor: [22, 94],
    popupAnchor: [-3, -76],
    shadowSize: [68, 95],
    shadowAnchor: [22, 94]
})

let marker;

let ws = new WebSocket("ws://localhost:8000/ws/location")

map.on("click", (e) => {
    ws.send(JSON.stringify({lat: e.latlng.lat, lng: e.latlng.lng}))
    console.log(e)
})

ws.onmessage = function (e) {
    const data = JSON.parse(e.data)
    console.log(data)

    if(marker) {
        map.removeLayer(marker)
    }
    marker = L.marker([data.lat, data.lng], {icon: mikroIcon}).addTo(map)
}

// navigator.geolocation.watchPosition(success, error)
//
// function success(pos) {
//     const lat = pos.coords.latitude
//     const lon = pos.coords.longitude
//
//     if(marker) {
//         map.removeLayer(marker)
//         map.removeLayer(circle)
//     }
//
//     marker = L.marker([lat, lon], {icon: mikroIcon}).addTo(map)
// }
//
// function error(err) {
//     console.log(err)
// }