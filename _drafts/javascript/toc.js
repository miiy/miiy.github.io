window.onload = function(){
    genPostDirectory();
};

/**
 * gen post directory
 * 
 * <div class="post-directory">
 *    <h2 style="font-size: 32px;">Table of Contents</h2>
 *    <ul style="margin:0 0 15px 0;">
 *      <li style="margin-left: {hn} * 20px;" data-hn="{hTag}">
 *        <a href="#dir{i}">{textContent}</a>
 *      </li>
 *    </ul>
 *  </div>
 * 
 * Jquery
 *
 * $(".post-content").before('<div class="post-directory"><h2 style="font-size: 32px;">Table of Contents</h2><ul style="margin:0 0 15px 0;"></ul></div>');
 * $(".post-content h1, \
 * .post-content h2, \
 * .post-content h3, \
 * .post-content h4, \
 * .post-content h5, \
 * .post-content h6").each(function(i, item){
 *     var localName = $(item)[0].localName;
 *    var hn = localName.match(/[\d]+/);
 *    // #dir[n] Avoid duplication
 *     $(item).attr("id", "dir" + i);
 *     $(".post-directory ul").append('<li style="margin-left:'+ hn * 20 +'px;" data-h="' + localName + '"><a href="#dir' + i + '">' + $(item).text() + '</a></li>');
 * });
 */
function genPostDirectory()
{
    var postContentEle = document.getElementsByClassName("post-content")[0];
    if (postContentEle === undefined) {
        return;
    }

    var ulEle = document.createElement("ul");
    ulEle.style.margin = "0 0 15px 0";

    for(i in postContentEle.childNodes) {
        var childEle = postContentEle.childNodes[i];
        var HTags = ["h1", "h2", "h3", "h4", "h5", "h6"];
        if (-1 !== HTags.indexOf(childEle.localName)) {
            var hn = childEle.localName.match(/[\d]+/);
            // #dir[n] Avoid duplication
            childEle.setAttribute("id", "dir" + i);

            var liEle = document.createElement("li");
            liEle.style.marginLeft = hn * 20 + "px";
            liEle.setAttribute("data-hn", childEle.localName);

            var aEle = document.createElement("a");
            aEle.href = "#dir" + i;
            aEle.innerText = childEle.textContent;
            
            liEle.appendChild(aEle);
            ulEle.appendChild(liEle);
        }
    }


    var h2Ele = document.createElement("h2");
    h2Ele.setAttribute("style", "font-size: 32px;");
    h2Ele.innerHTML = "Table of Contents";

    var postDirEle = document.createElement('div');
    postDirEle.className = "post-directory";
    postDirEle.appendChild(h2Ele);
    postDirEle.appendChild(ulEle);
    if (ulEle.children.length > 0) {
        postContentEle.insertBefore(postDirEle, postContentEle.children[0]);
    }
}


