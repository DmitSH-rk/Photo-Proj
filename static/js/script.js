document.addEventListener("DOMContentLoaded", function () {
  var swiper = new Swiper(".swiper", {
    loop: true,
    autoplay: {
      delay: 2000,
      disableOnInteraction: false,
    },
  });
});
// document.getElementById("btn-open").addEventListener("click", function () {
//   document.getElementsById("my-modal").classList.add("open");
// });

// document.getElementById("btn-close").addEventListener("click", function () {
//   document.getElementById("my-modal").classList.remove("open");
// });
document.getElementById("saveButton").addEventListener("click", function () {
  const name = document.getElementById("Name").value;
  const pass = document.getElementById("Password").value;

  fetch("/save", {
    method: "POST",
    headers: {
      "Content-Type": "application/x-www-form-urlencoded",
    },
    body: new URLSearchParams({
      Name: name,
      Password: pass,
    }),
  })
    .then((response) => response.text())
    .then((data) => {
      console.log("Success:", data);
      document.getElementById("Name").value = "";
      document.getElementById("Password").value = "";
    })
    .catch((error) => {
      console.error("Error:", error);
    });
});
// document.getElementById("loginButton").addEventListener("click", function () {
//   const name = document.getElementById("loginName").value;
//   const pass = document.getElementById("logPassword").value;

//   fetch("/check", {
//     method: "POST",
//     headers: {
//       "Content-Type": "application/x-www-form-urlencoded",
//     },
//     body: new URLSearchParams({
//       Name: name,
//       Password: pass,
//     }),
//   })
//     .then((response) => response.text())
//     .then((data) => {
//       console.log("Success:", data);
//       document.getElementById("loginName").value = "";
//       document.getElementById("logPassword").value = "";
//       // Обработка ответа
//       if (data === "User exists") {
//         alert("Login successful");
//         window.location.href = "/gayhouse";
//       } else {
//         alert("Invalid credentials");
//       }
//     })
//     .catch((error) => {
//       console.error("Error:", error);
//     });
// });
// document.addEventListener("DOMContentLoaded", function () {
//   const productImages = document.querySelectorAll(".main__img");
//   const imageUrls = document.querySelectorAll(".imageUrl");
//   productImages.forEach((img, index) => {
//     imageUrls[index].value = img.src;
//   });
// });
