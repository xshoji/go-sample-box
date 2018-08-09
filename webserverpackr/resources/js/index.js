function displayAlert(){
    alert("Clicked!")
}

var bodyElement = document.getElementById("body");
bodyElement.addEventListener("click", function(){
    displayAlert();
})
