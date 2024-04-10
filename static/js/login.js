

document.getElementById('loginForm').addEventListener('submit', function(event) {
    event.preventDefault();

    var formData = new FormData(event.target);
    console.log(formData);
    authLogin(formData);
})