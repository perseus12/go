$uscore="this_is_the_string_to_be_converted"
$arr=(${uscore//_/ })
printf %s "${arr[@]^}"
