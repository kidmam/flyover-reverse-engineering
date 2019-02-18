
This is an attempt to reverse-engineer the Flyover feature in Apple Maps. Main goal is to document the results and to provide code that emerges.

#### Motivation
Noticed differences between Google Earth and Apple Flyover during [previous project](https://github.com/retroplasma/earth-reverse-engineering). Extreme example:

<img src="https://user-images.githubusercontent.com/46618410/52183147-db89e500-27fc-11e9-9c75-fc78ff6cda58.jpg" alt="Google" title="Google"  width=50%><img src="https://user-images.githubusercontent.com/46618410/52183145-d62c9a80-27fc-11e9-9396-2d0acb34ec03.jpg" alt="Apple" title="Apple" width=50%>

#### General
Data in stored in map tiles:

|Type  | Purpose                                     | URL structure                                        |
|------|---------------------------------------------|------------------------------------------------------|
|DTM 1 | Terrain/Surface/Elevation                   | 🅐(?\|&)style=16&v=⓿&region=❶&x=❹&y=❺&z=❻         |
|DTM 2 | Terrain/Surface/Elevation                   | 🅐(?\|&)style=17&v=⓿&size=❷&scale=❸&x=❹&y=❺&z=❻  |
|C3M   | Texture, Mesh, ECEF, Rotation, Scene Graph? | 🅐(?\|&)style=15&v=⓿&region=❶&x=❹&y=❺&z=❻&h=❼    |
|C3MM  | Metadata                                    | 🅐(?\|&)style=52&v=⓿&region=❶&x=❹&y=❺&z=❻&h=❼    |   

- 🅐: URL prefix from resource manifest
- ⓿: Version from resource manifest or altitude manifest using region
- ❶: From altitude manifest
- ❷❸: todo
- ❹❺❻: Map tile. Regular scheme?
- ❼: Height/altitude? Only for C3M and C3MM



#### Authenticate URLs
```
go run cmd/auth/main.go [url] [session_id] [token_1] [token_2]
```
- `session_id`: 40 digits
- `token_1`: see GeoServices binary (function: GEOURLAuthenticationGenerateURL)
- `token_2`: see Geo Resource Manifest config (protobuf field 30)

#### Parse C3M file (WIP)
```
go run cmd/parse-c3m/main.go [file]
```

#### Files on macOS
- `~/Library/Preferences/com.apple.GEO.plist`
- `~/Library/Caches/GeoServices/Resources/altitude-*.xml`
- `~/Library/Containers/com.apple.geod/Data/Library/Caches/com.apple.geod/MapTiles/MapTiles.sqlitedb`
- `/System/Library/PrivateFrameworks/GeoServices.framework/GeoServices`
- `/System/Library/PrivateFrameworks/VectorKit.framework/VectorKit`
- `/System/Library/PrivateFrameworks/GeoServices.framework/XPCServices/com.apple.geod.xpc`

#### Important
THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
