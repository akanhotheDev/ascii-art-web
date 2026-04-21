function copyResult() {
    const text = document.querySelector("#result pre").innerText;

    navigator.clipboard.writeText(text)
        .then(() => {
            alert("Copied to clipboard!");
        })
        .catch(err => {
            alert("Failed to copy");
        });
}
function downloadResult() {
    const text = document.querySelector("#result pre").innerText;

    const blob = new Blob([text], { type: "text/plain" });
    const url = URL.createObjectURL(blob);

    const a = document.createElement("a");
    a.href = url;
    a.download = "ascii-art.txt";
    a.click();

    URL.revokeObjectURL(url);

}

src="https://unpkg.com/lucide@latest/dist/umd/lucide.min.js"