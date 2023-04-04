document.getElementById('creationDateFrom').addEventListener('input', function functionName(e) {
    let end = e.target.value;
    document.getElementById('creationDateFromNr').innerHTML = end;
});
document.getElementById('creationDateTo').addEventListener('input', function functionName(e) {
    let start = e.target.value;
    document.getElementById('creationDateToNr').innerHTML = start;
});

function showFilters() {
    const filters = document.getElementById("filters");
    const h2 = document.getElementById("filtersH2");

    console.log(h2);
    if (filters.style.display === "block") {
        filters.style.display = "none";
        h2.style.border = "1px solid var(--middle)";
    } else {
        filters.style.display = "block";
        h2.style.border = "none";
    }
}

let visible = "";

function showInfo(id) {
    document.getElementById("selection").style.display = "block";
    if (visible.length > 0) {
        document.getElementById(visible).style.display = "none";
        visible = "";
    }
    document.getElementById(id).style.display = "block";
    visible = id;
}

function removeHash() {
    location.hash = "";
}