(async () => {
  let blogs = await fetch("/api/blogs");
  blogs = await blogs.json();

  cHolder = document.getElementById("c-holder");

  blogs.Blogs.forEach((blog) => {
    let column = document.createElement("div");
    column.className = "col-md-4";
    column.id = "card-holder";

    let card = document.createElement("div");
    card.className = "card text-bg-dark";

    let cardImg = document.createElement("img");
    cardImg.src = blog.picture;
    cardImg.className = "card-img-top";

    let cardBody = document.createElement("div");
    cardBody.className = "card-body";

    let title = document.createElement("h5");
    title.innerText = blog.title;
    title.className = "card-title";

    let footer = document.createElement("div");
    let aDay = new Date(blog.date);
    footer.innerText = timeSince(aDay);
    footer.className = "card-footer";

    cardBody.appendChild(title);
    card.appendChild(cardImg);
    card.appendChild(cardBody);
    card.appendChild(footer);
    cHolder.appendChild(column);
    column.appendChild(card);

    document.getElementById("c-holder").append(column);
  });
})();

function timeSince(date) {
  var seconds = Math.floor((new Date() - date) / 1000);

  var interval = Math.floor(seconds / 31536000);

  if (interval > 1) {
    return interval + " years ago";
  }
  interval = Math.floor(seconds / 2592000);
  if (interval > 1) {
    return interval + " months ago";
  }
  interval = Math.floor(seconds / 86400);
  if (interval > 1) {
    return interval + " days ago";
  }
  interval = Math.floor(seconds / 3600);
  if (interval > 1) {
    return interval + " hours ago";
  }
  interval = Math.floor(seconds / 60);
  if (interval > 1) {
    return interval + " minutes ago";
  }
  return Math.floor(seconds) + " seconds ago";
}
