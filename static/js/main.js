$(document).ready(function(){
    $(".post-page .post-content h2,h3,h4,h5,h6").each(function(i,item){
        var tag = $(item).get(0).localName;
        $(item).attr("id","dir"+i);
        $(".post-directory").append('<a class="new'+tag+'" href="#dir'+i+'">'+$(this).text()+'</a>');
        $(".newh2").css("padding-left",0);
        $(".newh3").css("padding-left",25);
        $(".newh4").css("padding-left",50);
        $(".newh5").css("padding-left",75);
        $(".newh6").css("padding-left",100);
    });
    $('.post-directory').affix({
        offset: {
            top: function () {
                var offsetTop = $('.post-directory').offset().top;
                return (this.top = offsetTop)
            },
            bottom: function() {
                return (this.bottom = $('.footer').outerHeight(true));
            }
        }
    });
});


particlesJS('particles', 
  {
    "particles": {
      "number": {
        "value": 20,
        "density": {
          "enable": true,
          "value_area": 1000
        }
      },
      "color": {
        "value": "#e1e1e1"
      },
      "shape": {
        "type": "circle",
        "stroke": {
          "width": 0,
          "color": "#000000"
        },
        "polygon": {
          "nb_sides": 5
        },
        "image": {
          "src": "img/github.svg",
          "width": 100,
          "height": 100
        }
      },
      "opacity": {
        "value": 0.5,
        "random": false,
        "anim": {
          "enable": false,
          "speed": 1,
          "opacity_min": 0.1,
          "sync": false
        }
      },
      "size": {
        "value": 15,
        "random": false,
        "anim": {
          "enable": false,
          "speed": 180,
          "size_min": 0.1,
          "sync": false
        }
      },
      "line_linked": {
        "enable": true,
        "distance": 650,
        "color": "#cfcfcf",
        "opacity": 0.26,
        "width": 1
      },
      "move": {
        "enable": true,
        "speed": 2,
        "direction": "none",
        "random": true,
        "straight": false,
        "out_mode": "out",
        "attract": {
          "enable": false,
          "rotateX": 600,
          "rotateY": 1200
        }
      }
    },
    "interactivity": {
      "detect_on": "canvas",
      "events": {
        "onhover": {
          "enable": false,
          "mode": "repulse"
        },
        "onclick": {
          "enable": false,
          "mode": "push"
        },
        "resize": true
      },
      "modes": {
        "grab": {
          "distance": 400,
          "line_linked": {
            "opacity": 1
          }
        },
        "bubble": {
          "distance": 400,
          "size": 40,
          "duration": 2,
          "opacity": 8,
          "speed": 3
        },
        "repulse": {
          "distance": 200
        },
        "push": {
          "particles_nb": 4
        },
        "remove": {
          "particles_nb": 2
        }
      }
    },
    "retina_detect": true,
    "config_demo": {
      "hide_card": false,
      "background_color": "#b61924",
      "background_image": "",
      "background_position": "50% 50%",
      "background_repeat": "no-repeat",
      "background_size": "cover"
    }
  }

);
