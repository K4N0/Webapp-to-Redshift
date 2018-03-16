var slider_content = document.getElementById('box');
   
   //contain image in an array
   var image = ['a', 'b', 'c', 'd', 'e' ];

   var i = image.length;

   //function for next slide

   function nextImage(){
     if (i < image.length) {
       i= i+1;
     }else{
        i = 1;
     }
       slider_content.inner_HTML = "<img src"+image[i-1]+".jpg>";
   }

   //function for prev slide

   function prevImage() {

    if (i < image.length +1 && i > 1) {
       i = i-1;
     } else {
        i = image.length;
     }
       slider_content.inner_HTML = "< img src" + image[i-1] + ".jpg>";
}

var modal = document.getElementById('myModal');
var img = document.getElementById('myImg');
var modalImg = document.getElementById("img01");
var captionText = document.getElementById("caption");
img.onclick = function(){
    modal.style.display = "block";
    modalImg.src = this.src;
    captionText.innerHTML = this.alt;
}

var span = document.getElementsByClassName("close")[0];
  
span.onclick = function() { 
    modal.style.display = "none";
}