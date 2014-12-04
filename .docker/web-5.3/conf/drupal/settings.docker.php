<?php
/**
 * @file
 * Docker specific configuration.
 */

// Drupal 6 connection string.
$db_url = 'mysql://dbuser:dbpass@dbhost/drupal';
// Set the host to the proxied container IP.
// $base_url = 'http://projectname';
// Set file system paths.
$conf['file_public_path'] = 'sites/default/files';
$conf['file_temporary_path'] = '/tmp';

// Set files-private to be one level up from the docroot.
$conf['file_directory_path'] = __DIR__ . '/../../files-private';

// Ensure mail is local only. (My sandbox mail server is safe.)
$conf['smtp_library'] = 'sites/all/modules/contrib/devel/devel.module';
$conf['site_mail'] = 'admin@localhost';

// Make sure logging is enabled and at a reasonable level.
// (This could probably use some tweaking...)
error_reporting(E_ALL & ~E_STRICT & ~E_NOTICE & ~E_DEPRECATED);
$conf['error_level'] = 0;
ini_set('display_errors', FALSE);

// Misc sandbox specific settings.
$conf['apachesolr_read_only'] = "1";
$conf['gtranslate_api_key'] = 'AIzaSyBxd3lZI9Z-LHlkXuoeKBz2fBgREGxJ5uY';
$conf['site_mail'] = 'admin@localhost';
$conf['securepages_enable'] = 0;
