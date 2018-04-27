package droictx

const (
	HTTPHeaderAppID            = "X-Droi-AppID"
	HTTPHeaderAppIDMode        = "X-Droi-AidMode"
	HTTPHeaderDeviceID         = "X-Droi-DeviceID"
	HTTPHeaderUserID           = "X-Droi-UserID"
	HTTPHeaderRequestID        = "X-Droi-ReqID"
	HTTPHeaderPlatformKey      = "X-Droi-Platform-Key"
	HTTPHeaderAPIKey           = "X-Droi-Api-Key"
	HTTPHeaderServiceAppID     = "X-Droi-Service-AppID"
	HTTPHeaderServiceAppIDMode = "X-Droi-SAidMode"
	HTTPContextKey             = "X-Ctx"
	HTTPHeaderRole             = "X-Droi-Role"
	HTTPHeaderSessionToken     = "X-Droi-Session-Token"
	HTTPHeaderURI              = "X-Droi-URI"
	HTTPHeaderMethod           = "X-Droi-Method"
	HTTPHeaderRemoteIP         = "X-Droi-Remote-IP"
	HTTPHeaderRemotePort       = "X-Droi-Remote-Port"
	HTTPHeaderSlotID           = "X-Droi-SlotID"
	HTTPHeaderServiceAppCheat  = "X-Droi-Service-AppCheat"
	ShortAppID                 = "Aid"
	ShortAppIDMode             = "Aidm"
	ShortDeviceID              = "Did"
	ShortUserID                = "Uid"
	ShortRequestID             = "Rid"
	ShortPlatformKey           = "XPk"
	ShortAPIKey                = "Ak"
	ShortServiceAppID          = "SAid"
	ShortServiceAppIDMode      = "SAidm"
	ShortRole                  = "R"
	ShortSessionToken          = "St"
	ShortURI                   = "XUri"
	ShortMethod                = "XMd"
	ShortRemoteIP              = "XIp"
	ShortRemotePort            = "XPort"
	ShortSlotID                = "Slid"
	ShortServiceAppCheat       = "SAc"
	// this is only used in GoBuster and Accelerator for Push UDP
	ShortSessionID = "Sid"

	//The most important Key !
	SystemKey = "2BMustDie"
)

//This interface designed for getting DroiCtx from fasthttp *RequestHeader
type Peeker interface {
	Peek(key string) []byte
}

//This interface designed for setting fasthttp *RequestHeader, net/http Header with DroiCtx
type Setter interface {
	Set(key, value string)
}

//Getter interface designed for getting DroiCtx from net/http Header
type Getter interface {
	Get(key string) string
}


var (
	sKMap, hKMap map[string]string
)

func init() {
	sKMap = map[string]string{
		ShortAppID:            HTTPHeaderAppID,
		ShortAppIDMode:        HTTPHeaderAppIDMode,
		ShortDeviceID:         HTTPHeaderDeviceID,
		ShortUserID:           HTTPHeaderUserID,
		ShortRequestID:        HTTPHeaderRequestID,
		ShortPlatformKey:      HTTPHeaderPlatformKey,
		ShortAPIKey:           HTTPHeaderAPIKey,
		ShortServiceAppID:     HTTPHeaderServiceAppID,
		ShortServiceAppIDMode: HTTPHeaderServiceAppIDMode,
		ShortRole:             HTTPHeaderRole,
		ShortSessionToken:     HTTPHeaderSessionToken,
		ShortURI:              HTTPHeaderURI,
		ShortMethod:           HTTPHeaderMethod,
		ShortRemoteIP:         HTTPHeaderRemoteIP,
		ShortRemotePort:       HTTPHeaderRemotePort,
		ShortSlotID:           HTTPHeaderSlotID,
		ShortServiceAppCheat:  HTTPHeaderServiceAppCheat,
	}

	hKMap = map[string]string{
		HTTPHeaderAppID:            ShortAppID,
		HTTPHeaderAppIDMode:        ShortAppIDMode,
		HTTPHeaderDeviceID:         ShortDeviceID,
		HTTPHeaderUserID:           ShortUserID,
		HTTPHeaderRequestID:        ShortRequestID,
		HTTPHeaderPlatformKey:      ShortPlatformKey,
		HTTPHeaderAPIKey:           ShortAPIKey,
		HTTPHeaderServiceAppID:     ShortServiceAppID,
		HTTPHeaderServiceAppIDMode: ShortServiceAppIDMode,
		HTTPHeaderRole:             ShortRole,
		HTTPHeaderSessionToken:     ShortSessionToken,
		HTTPHeaderURI:              ShortURI,
		HTTPHeaderMethod:           ShortMethod,
		HTTPHeaderRemoteIP:         ShortRemoteIP,
		HTTPHeaderRemotePort:       ShortRemotePort,
		HTTPHeaderSlotID:           ShortSlotID,
		HTTPHeaderServiceAppCheat:  ShortServiceAppCheat,
	}
}

// retrun a map, key is HTTP Header Field, value is Short Field
func IFieldHeaderKeyMap() (keyMap map[string]string) {
	return hKMap
}

// retrun a map, key is Short Field, value is HTTP Header Field
func IFieldShortKeyMap() (keyMap map[string]string) {
	return sKMap
}

func GetContextFromPeeker(p Peeker) Context {
	c := Context{}
	var v []byte
	for hk, sk := range hKMap {
		v = p.Peek(hk)
		if len(v) > 0 {
			c.Set(sk, string(v))
		}
	}
	return c
}

func GetContextFromGetter(g Getter) Context {
	c := Context{}
	var v string
	for hk, sk := range hKMap {
		v = g.Get(hk)
		if len(v) > 0 {
			c.Set(sk, v)
		}
	}
	return c
}

func (c *Context) SetHTTPHeaders(s Setter) {
	for hk, sk := range c.HeaderMap() {
		s.Set(hk, sk)
	}
}

// retrun a map, key HTTP Header Field, value is the field value stored in Context
func (c *Context) HeaderMap() (ret map[string]string) {
	ret = make(map[string]string)
	for sk, hk := range sKMap {
		v, _ := c.GetString(sk)
		if len(v) > 0 {
			ret[hk] = v
		}
	}
	return
}

// Set Context with Header Key, Store in Context with Short Key
func (c *Context) HeaderSet(headerField, headerValue string) {
	if sk, ok := hKMap[headerField]; ok {
		c.Set(sk, headerValue)
	}
	return
}
