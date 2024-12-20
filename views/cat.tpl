<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Cat Page</title>
    <link href="https://cdn.jsdelivr.net/npm/remixicon@4.5.0/fonts/remixicon.css" rel="stylesheet"/>
    <link rel="stylesheet" href="static/css/cat.css">
    <script src="https://code.jquery.com/jquery-3.6.0.min.js"></script>
</head>
<body>
    <div class="main-content">
        <nav style="display: flex; gap: 30px;">
            <div style="text-align: center;" onclick="showVoting()">
                <i class="ri-arrow-up-down-line"></i>
                <p id="voting" class="nav-item">Voting</p> 
            </div>
            <div style="text-align: center;">
                <i class="ri-search-eye-line"></i>
                <p id="breads" class="nav-item">Breads</p>
            </div>
            <div style="text-align: center;" onclick="showFavorites()">
                <i class="ri-heart-2-line"></i>
                <p id="favs" class="nav-item">Favs</p>
            </div>
        </nav>

        <!-- Voting Section -->
        <div id="randomCatImageSection" style="display: block;">
            <div id="loading" class="loading-animation">
                <img src="static/img/cat.png"/>
            </div>
            <img id="catImage" src="{{.CatImageURL}}" alt="Random Cat" style="display: none;">
        </div>

        <!-- Favorites Section -->
        <div id="fav-container"  style="display: none;">
            <div class="view-options">
                <button class="active" id="grid-view" onclick="switchLayout('grid')">
                    <i class="ri-layout-grid-fill"></i> 
                </button>
                <button id="scroll-view" onclick="switchLayout('scroll')">
                    <i class="ri-layout-horizontal-line"></i> 
                </button>
            </div>
            <div id="fav-image-container" class="grid-container">
            
            </div>
        </div>

        <div class="footer" id="footerSection">
            <div class="heart-container">
                <i class="ri-heart-2-line" id="heart-icon" onclick="addToFavorites('{{.CatImageID}}')"></i>
                <span class="heart-animation"></span>
            </div>
            <div class="thumbs" style="gap: 10px;">
                <i class="ri-thumb-up-line" id="thumb-up" onclick="sendVote('{{.CatImageID}}', true)"></i>
                <i class="ri-thumb-down-line" id="thumb-down" onclick="sendVote('{{.CatImageID}}', false)"></i>
            </div>
        </div>
    </div>

    <script>
        const apiKey = "{{.API_Key}}";
    </script>
    <script src="static/js/loading.js"></script>
    <script src="static/js/voting.js"></script>
    <script src="static/js/addToFav.js"></script>
</body>
</html>
