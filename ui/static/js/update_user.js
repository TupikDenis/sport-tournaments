$('#update_user_password').submit(function(e){
        e.preventDefault();
        var formData = $(this).serialize();
        var formAction = $(this).attr('action')
        $.ajax({
            url: formAction,
            data: formData,
            type: "PATCH",

            success: function(data){
                window.location.href = "/profile";
            }
        });
});


