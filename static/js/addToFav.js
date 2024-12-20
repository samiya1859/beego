function addToFavorites(imageId) {
   
    const data = {
        image_id: imageId,
        sub_id: 'my-user-1234'  
    };

   
    fetch('https://api.thecatapi.com/v1/favourites', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
            'x-api-key': apiKey 
        },
        body: JSON.stringify(data)  
    })
    .then(response => {
        if (response.ok) {
            alert('Cat added to favorites!');
            console.log("Cat added to the favs!!!!!")
        } else {
            alert('Failed to add cat to favorites.');
        }
    })
    .catch(error => {
        console.error('Error:', error);
        alert('Failed to add cat to favorites.');
    });
}


// get all favs
function showVoting() {
    document.getElementById("randomCatImageSection").style.display = "block";
    document.getElementById("fav-container").style.display = "none";
    document.getElementById("footerSection").style.display = "flex";
}

function showFavorites() {
    document.getElementById("randomCatImageSection").style.display = "none";
    document.getElementById("fav-container").style.display = "block";
    document.getElementById("footerSection").style.display = "none";
    loadFavorites();  // Fetch and display the favorites images
}


// Function to toggle between Grid and Scroll layout
function switchLayout(layout) {
    const favImageContainer = document.getElementById('fav-image-container');
    
    
    favImageContainer.classList.remove('grid-container', 'scroll-container');

    
    if (layout === 'grid') {
        favImageContainer.classList.add('grid-container');
    } else if (layout === 'scroll') {
        favImageContainer.classList.add('scroll-container');
    }
}


function loadFavorites() {
    fetch('https://api.thecatapi.com/v1/favourites', {
        method: 'GET',
        headers: {
            'x-api-key': apiKey  
        }
    })
    .then(response => response.json())
    .then(data => {
        const favContainer = document.getElementById("fav-image-container");
        favContainer.innerHTML = ""; 

        data.forEach(fav => {
            const img = document.createElement("img");
            img.src = fav.image.url;  
            img.alt = "Favorite Cat Image";
            img.style.width = "150px"; 
            img.style.margin = "10px";
            img.style.borderRadius = "8px";  
            favContainer.appendChild(img); 
        });
    })
    .catch(error => {
        console.error('Error fetching favorites:', error);
    });

    switchLayout('grid');
}

// Initially show Voting section
showVoting();