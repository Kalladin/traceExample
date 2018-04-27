# v0.5.0
- errors
    - Add TraceDroiError for recording tracable error stacks
    - Remove NewApiError
- rdb
    - Migrate NewApiError to ConstDroiError

# v0.4.3
- rdb
    - Refactor rdb struct & error
    - Add QueryAppPrefixPayload for query app datas with appID prefix

# v0.4.2
- rdb
    - Add rdb struct & error into droipkg

# v0.4.1
- file
    - Add ParseCdnURL() to get fqdn, qppid, fid, and filename from CDN url

# v0.4.0
- file
    - Modify file.GetFileBasename CDN link rule

# v0.3.0
- file
    - Change the Op field of majesty's CallBackPayload to map[string]interface{}
    - Add log fields
    - Add OpCodes ID and mapping

# v0.2.0
- file
    - Add file.GetFileBasename to get basename for CDN link

# v0.1.0
- init. commit
