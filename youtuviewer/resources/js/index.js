var formElement = document.getElementById("form-container");
// formElement.addEventListener("click", function(){
//     var json = await getJson("test") // call go function
//     alert(json)
// })
formElement.addEventListener("submit", function(event) {
    event.preventDefault();
    return false;
});

formElement.addEventListener("submit", async () => {
    var name = document.getElementById("form-name").value;
    var tel = document.getElementById("form-tel").value;
    var email = document.getElementById("form-email").value;
    var json = await getJson(name, tel, email); // call go function
    alert(json);
});
