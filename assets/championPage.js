function searchChampion() {
    var input = document.getElementById("champSearch").value.trim();
    if (input) {
        window.location.href = "/champions/" + encodeURIComponent(input.toLowerCase());
    } else {
        alert("Please enter a champion name");
    }
}
