select profile.full_name from profile left join profile_manager 
on profile.profile_id=profile_manager.profile_id where
profile_manager.user_account_id=$1 limit 10 offset 0;


select count(*) from profile left join profile_manager 
on profile.profile_id=profile_manager.profile_id where
profile_manager.user_account_id=$1;



