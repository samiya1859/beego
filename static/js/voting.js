function sendVote(imageId, isUpvote) {
   
    var voteValue = isUpvote ? 1 : 2;  // 1 for upvote, 2 for downvote

    // Create the request payload
    var payload = {
        image_id: imageId,
        sub_id: "my-user-1234", 
        value: voteValue
    };

    // Send the POST request using fetch
    fetch("https://api.thecatapi.com/v1/votes", {
        method: "POST",
        headers: {
            "Content-Type": "application/json" ,
            "x-api-key":apiKey
            
        },
        body: JSON.stringify(payload)  
    })
    .then(response => response.json())  
    .then(data => {
        console.log("Vote successful", data);
        alert("Vote successful");
    })
    .catch(error => {
        console.error("Vote failed: ", error);
        alert("Vote failed: " + error.message);  
    });
}
