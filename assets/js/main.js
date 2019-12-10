$(document).ready(function(){
    $(".post-content").before('<div class="post-directory"><h2>Table of Contents</h2><ul style="margin:0 0 15px 0;"></ul></div>');
    $(".post-content h1, \
    .post-content h2, \
    .post-content h3, \
    .post-content h4, \
    .post-content h5, \
    .post-content h6").each(function(i, item){
        var localName = $(item)[0].localName;
        var hn = localName.match(/[\d]+/);
        // #dir[n] Avoid duplication
        $(item).attr("id", "dir" + i);
        $(".post-directory ul").append('<li style="margin-left:'+ hn * 20 +'px;" data-h="' + localName + '"><a href="#dir' + i + '">' + $(item).text() + '</a></li>');
    });
});
