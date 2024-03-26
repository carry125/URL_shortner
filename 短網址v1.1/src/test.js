document.addEventListener("DOMContentLoaded", function () {
    const queryForm = document.getElementById("queryForm");
    const resultDiv = document.getElementById("result");

    queryForm.addEventListener("submit", function (e) {
        e.preventDefault();
        
        const formData = new FormData();
        formData.append("url", document.getElementById("start").value); //append("url") 此為API的PostForm指定字

        // 發送API
        fetch("http://localhost:8080/shortner", {
            method: "POST",
            body: formData
        })
        .then(response => response.json())
        .then(data => {  //data傳json出來 所以data底下的資料都會是json模樣
                if (data.error) {
                    resultDiv.innerHTML = `<p>Error: ${data.error}</p>`;
                } else {
                    // 显示短網址
                    resultDiv.innerHTML = `<h2>短網址：</h2>`;
                    const shortnerCode = data.code; //golang "code"
                    resultDiv.innerHTML += `<p> http://localhost:8080/shortner/${shortnerCode}</p>`;
                }
            })
            .catch(error => {
                resultDiv.innerHTML = `<p>Error: ${error.message}</p>`;
            });
    });
});