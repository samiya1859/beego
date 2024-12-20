const img = document.getElementById('catImage');
const loading = document.getElementById('loading');

img.onload = function() {
    loading.style.display = 'none';
    img.style.display = 'block'; 
};


const navItems = document.querySelectorAll('.nav-item');

   console.log(navItems)
    navItems.forEach(item => {
        item.addEventListener('click', function() {
           
            navItems.forEach(i => i.classList.remove('active'));
            this.classList.add('active');
        });
    });




    