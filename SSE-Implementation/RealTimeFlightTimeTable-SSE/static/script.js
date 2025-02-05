const eventSource = new EventSource("/flights");
eventSource.onmessage = function(event) {
    const flight = JSON.parse(event.data);
    const table = document.getElementById("flightTable");

    let row = document.createElement("tr");
    row.innerHTML = `<td>${flight.flightNo}</td><td>${flight.status}</td><td>${flight.eta}</td>`;

    table.insertBefore(row, table.firstChild);
};

eventSource.onerror = function() {
    console.error("SSE connection lost");
};