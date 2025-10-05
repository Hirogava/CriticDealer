const reqUrl = `http://localhost:8080/api/v1/critical`;
const key = '9abec63a-4211-4ea4-94ae-8a0d41c26d81'

const modal = ` <div class="bottom-sheet" id="warningSheet">
  <div class="handle"></div>
  <button class="close-btn" onclick="document.getElementById('warningSheet').remove()">×</button>
  <div class="warning">
  <img src="http://localhost:8080/images/dangersmall.png" alt="" class="warn-icon">   
    
    Будьте осторожны!
  </div>
  <p>На данном маршруте при аналогичных погодных условиях зафиксировано большое количество ДТП.</p>
  <p>Начинающим водителям рекомендуем выбрать более безопасный маршрут.</p>
</div>`;

const map = new mapgl.Map('container', {
    center: [37.615655, 55.768005],
    zoom: 13,
    key: key,
});

const directions = new mapgl.Directions(map, {
    directionsApiKey: key,
});

const markers = [];
let isDangerous = false;
const critical_markers = [];

let firstPoint;
let secondPoint;
let selecting = 'a';
const controlsHtml = `<button id="reset">Reset points</button> `;
new mapgl.Control(map, controlsHtml, {
    position: 'topLeft',
});
const resetButton = document.getElementById('reset');
resetButton.addEventListener('click', function() {
    selecting = 'a';
    firstPoint = undefined;
    secondPoint = undefined;
    if (routeLine) {
        routeLine.destroy();
    }

    if (markers && markers.length) {
        markers.forEach(m => {
            try { m.destroy(); } catch (e) { console.warn('Failed to destroy marker', e); }
        });

        critical_markers.forEach(m => {
            try { m.destroy(); } catch (e) { console.warn('Failed to destroy marker', e); }
        });
        markers.length = 0;
        isDangerous = false;
    }
});

map.on('click', (e) => {
    const coords = e.lngLat;
    if (selecting != 'end') {
        markers.push(
                new mapgl.Marker(map, {
                    coordinates: coords,
                    icon: 'https://docs.2gis.com/img/dotMarker.svg',
                }),
        );
    }
    if (selecting === 'a') {
        firstPoint = {
            lon:coords[0],
            lat:coords[1]
        }
        selecting = 'b';
    } else if (selecting === 'b') {
        secondPoint = {
            type:"stop",
            lon:coords[0],
            lat:coords[1]
        }
        selecting = 'end';
    }
    if (firstPoint && secondPoint) {
        fetchRoute([firstPoint, secondPoint]);
    }
    });

function getCoordinates(outcoming_path){
    const selection = outcoming_path.geometry[0].selection;

    const coords = selection
    .match(/\d+\.\d+ \d+\.\d+/g)
    .map(pair => pair.split(' ').map(Number));

    const middle = coords[Math.floor(coords.length / 2)];
    const [lon, lat] = middle;

    return [lon, lat]

}

function renderCriticalPoints(maneuvers) {
    maneuvers.forEach((maneuver) => {
        if (maneuver.critical && maneuver.comment != "start") {
            isDangerous = true;
            critical_markers.push(
                new mapgl.Marker(map, {
                    coordinates: getCoordinates(maneuver.outcoming_path),
                    icon: 'http://localhost:8080/images/danger2gisorange.svg',
                    size: [25, 32],
                    anchor: [16, 32],
                })
            );
        }
    });
}

let routeLine;
function fetchRoute(points) {
    const body = JSON.stringify(
        {
            points: points,
            transport:"driving",
            output: "detailed",
            locale: "ru",
            alternative : 3
        }
    )
    console.warn(body);
    fetch(reqUrl, {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: body
    })
        .then((res) => {
            if (!res.ok) {
                throw new Error(`HTTP error! Status: ${res.status}`);
            }
            return res.json();
        })
        .then((parsed) => {
            console.log('fetchRoute parsed response:', parsed);
            if (parsed.result && parsed.result.length > 0) {
                const coordinates = parsed.result[0].maneuvers
                    .flatMap((maneuver) => {
                        if (
                            maneuver.outcoming_path &&
                            maneuver.outcoming_path.geometry &&
                            maneuver.outcoming_path.geometry.length > 0
                        ) {
                            return maneuver.outcoming_path.geometry
                                .flatMap((geometry) => {
                                    const selection = geometry.selection;
                                    return selection
                                        .replace("LINESTRING(", "")
                                        .replace(")", "")
                                        .split(",")
                                        .map((point) => point.trim().split(" ").map(Number));
                                });
                        }
                        return [];
                    });
                if (coordinates.length > 0) {
                    renderRoute(coordinates);
                    renderCriticalPoints(parsed.result[0].maneuvers);
                    if (isDangerous){
                        new mapgl.Control(map, modal, {
                            position: "bottomCenter"
                        });
                    }
                } else {
                    console.error("No coordinates found in response");
                }
            } else {
                console.error("No route found in response");
            }
        })
        .catch((err) => console.error("Error fetching route data:", err.message || err));
}

function renderRoute(coordinates) {
    if (routeLine) {
        routeLine.destroy();
    }
    routeLine = new mapgl.Polyline(map, {
        coordinates,
        width: 6,
        color: "#0078FF",
    });
}

// new mapgl.Control(map, modal, {
//                             position: "bottomCenter"
//                         });


// const marker = new mapgl.Marker(map, {
//     coordinates: [37.582591, 55.7],
//     // use a local relative path to avoid CORS when testing locally over HTTP
//     // make sure `danger2gisorange.png` is in the same `frontend/` folder
//     icon: './danger2gisorange.svg',
//     size: [25, 32],
//     anchor: [16, 32],
// });
