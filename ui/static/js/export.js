function download_file(text, file_name) {
    let blob = new Blob([text], { type: 'text/plain' });
    let link = document.createElement("a");
    link.download = file_name;
    //link.innerHTML = "Download File";
    link.href = window.URL.createObjectURL(blob);
    document.body.appendChild(link);
    link.click();
    setTimeout(() => {
        document.body.removeChild(link);
        window.URL.revokeObjectURL(link.href);
    }, 100);
}


function exportData(){
    var exp = document.getElementById("type")
    var value = exp.value;
    var text

    switch (value) {
        case "html":
            var html = document.getElementById("table").innerHTML
            text = html
            break
        case "json":
            var json = document.getElementById("json").textContent
            text = json
            break
        case "txt":
            var txt = document.getElementById("table").textContent
            text = txt
            break
    }

    download_file(text.trim(), document.getElementById("tournament_name").textContent.toLowerCase().replace(/\s+/g, '') + "." + value)
}