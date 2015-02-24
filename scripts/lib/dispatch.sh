# Changes zsh globbing patterns
unsetopt NO_MATCH >/dev/null 2>&1 || :

# Dispatches calls of commands and arguments
dispatch ()
{
	namespace="$1"     # Namespace to be dispatched
	arg="${2:-}"       # First argument
	short="${arg#*-}"  # First argument without trailing -
	long="${short#*-}" # First argument without trailing --

	# Exit and warn if no first argument is found
	if [ -z "$arg" ]; then
		"${namespace}_" # Call empty call placeholder
		return 1
	fi

	shift 2 # Remove namespace and first argument from $@

	# Detects if a command, --long or -short option was called
	if [ "$arg" = "--$long" ];then
		longname="${long%%=*}" # Long argument before the first = sign

		# Detects if the --long=option has = sign
		if [ "$long" != "$longname" ]; then
			longval="${long#*=}"
			long="$longname"
			set -- "$longval" "${@:-}"
		fi

		main_call=${namespace}_option_${long}


	elif [ "$arg" = "-$short" ];then
		main_call=${namespace}_option_${short}
	else
		main_call=${namespace}_command_${long}
	fi

	$main_call "${@:-}" && dispatch_returned=$? || dispatch_returned=$?

	if [ $dispatch_returned = 127 ]; then
		"${namespace}_call_" "$namespace" "$arg" # Empty placeholder
		return 1
	fi

	return $dispatch_returned
}
