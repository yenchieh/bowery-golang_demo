
$(function(){
    var $input = $('.jsSearchInput');
    $('.jsSearch').on("click", function(e) {
        e.preventDefault();
        search($input.val());
    });

    function search(query) {
        window.location.href="/search?q=" + query;
    }
});

