package test

import (
	"fmt"
	"github.com/Tony-Kwan/go-geo/geo"
	"github.com/Tony-Kwan/go-geo/geo/io/geojson"
	wkt2 "github.com/Tony-Kwan/go-geo/geo/io/wkt"
	"github.com/atotto/clipboard"
	"math"
	"testing"
)

var calc = geo.VectorCalculator{}

func TestPolygon_Triangulate(t *testing.T) {
	//wktStr := "POLYGON((-81.17061893429998 29.480056431500003,-81.18722876968559 29.469692395623255,-81.16458972779999 29.46415400310002,-81.1272890215 29.460051243099997,-81.133653453 29.4403458341,-81.16794171790002 29.45376429799998,-81.20013422968789 29.461639778161086,-81.20972296637115 29.455656695927516,-81.2058918196 29.429952449900014,-81.1955997485 29.398468924200003,-81.2189023361 29.395833882800005,-81.2181794089 29.4285629699,-81.22115672199574 29.44852237757047,-81.2410056151 29.449413482699995,-81.22296854975409 29.460668567001903,-81.22398374296428 29.46747425202564,-81.2442545563 29.472433243399976,-81.2815606519 29.47653024280001,-81.2752052098 29.496236469699994,-81.2409002435 29.4828225312,-81.22571922956847 29.479108663162705,-81.2273819469 29.49025524130002,-81.2376861542 29.52173806820001,-81.21436262759998 29.524373685799986,-81.21508686530001 29.491644609199994,-81.21274535382314 29.475934747199815,-81.20251445653352 29.47343187073794,-81.1778413341 29.488827838299983,-81.15084304890001 29.511618865999992,-81.13714804159999 29.494980623199993,-81.17061893429998 29.480056431500003))"
	//wktStr := "POLYGON((-119.14159855412011 41.92499096272812,-102.87886464247828 26.767453854346442,-90.9950327117614 32.905865438530256,-93.78787570746942 36.051610613180415,-98.38245074916146 33.171901863337226,-102.05698491988139 36.07009780383788,-96.33267709169169 38.664513680653954,-88.4319360597416 37.10206976053486,-87.28074151678913 30.775914227484662,-87.59809404666174 26.844389706437113,-79.64598553204937 29.904465341243238,-80.44042235294994 37.59795124939845,-84.81650967729065 43.819424657392034,-96.2837724335406 40.82461783417634,-98.73181999767975 46.03621512184026,-103.30774793950121 42.50917847188873,-108.37624078356396 46.00372124466807,-119.14159855412011 41.92499096272812))"
	wktStr := "POLYGON( (8.62883106519 50.0445149656, 8.534466861444404 50.02180734290968, 8.53497888571 49.9715565954, 8.51820004288 49.9714847, 8.517712187759079 50.01777553006917, 8.49779800804 50.0129834197, 8.49191126524 50.0230906883, 8.494910001894636 50.023812319904565, 8.49233921324 50.0282606405, 8.51753828493684 50.034276547784295, 8.517518586614017 50.036145651285224, 8.46067995035 50.0224946542, 8.45479960543 50.0326037399, 8.51739746573664 50.04763837876119, 8.51725514352 50.0611428256, 8.534065325 50.0612146676, 8.534162630197835 50.05166500510094, 8.570129366 50.0603034295, 8.57602082208 50.0501961968, 8.534279745046412 50.04017119693419, 8.53429903627058 50.0382779303452, 8.62344132599 50.0595593393, 8.62930301272 50.0494449958, 8.62636176421957 50.0487428364738, 8.62883106519 50.0445149656))"
	var polygon = wkt2.MustPolygon(wkt2.WktReader{}.Read(wktStr))
	tris, err := polygon.Triangulate()
	if err != nil {
		t.Error(err)
		return
	}
	ps := make([]geo.Polygon, len(tris))
	for i, tri := range tris {
		ps[i] = tri.ToPolygon()
	}
	clipboard.WriteAll(geojson.GeojsonWriter{}.EncodePolygons(ps))
}

func TestPolygon_ConvexHull(t *testing.T) {
	var wktStr = "POLYGON((-119.14159855412011 41.92499096272812,-102.87886464247828 26.767453854346442,-90.9950327117614 32.905865438530256,-93.78787570746942 36.051610613180415,-98.38245074916146 33.171901863337226,-102.05698491988139 36.07009780383788,-96.33267709169169 38.664513680653954,-88.4319360597416 37.10206976053486,-87.28074151678913 30.775914227484662,-87.59809404666174 26.844389706437113,-79.64598553204937 29.904465341243238,-80.44042235294994 37.59795124939845,-84.81650967729065 43.819424657392034,-96.2837724335406 40.82461783417634,-98.73181999767975 46.03621512184026,-103.30774793950121 42.50917847188873,-108.37624078356396 46.00372124466807,-119.14159855412011 41.92499096272812))"
	var polygon = wkt2.MustPolygon(wkt2.WktReader{}.Read(wktStr))
	hull, _ := polygon.ConvexHull()
	t.Log(hull.String())

	wkt := fmt.Sprintf("GEOMETRYCOLLECTION(%s,%s)", polygon.String(), hull.String())
	t.Log(wkt)
	clipboard.WriteAll(wkt)
}

func TestPolygon_Split(t *testing.T) {
	//wktStr := "POLYGON((-81.17061893429998 29.480056431500003,-81.18722876968559 29.469692395623255,-81.16458972779999 29.46415400310002,-81.1272890215 29.460051243099997,-81.133653453 29.4403458341,-81.16794171790002 29.45376429799998,-81.20013422968789 29.461639778161086,-81.20972296637115 29.455656695927516,-81.2058918196 29.429952449900014,-81.1955997485 29.398468924200003,-81.2189023361 29.395833882800005,-81.2181794089 29.4285629699,-81.22115672199574 29.44852237757047,-81.233787128 29.440641385499973,-81.260763335 29.417847494400007,-81.27445843 29.43448412120003,-81.2410056151 29.449413482699995,-81.22296854975409 29.460668567001903,-81.22398374296428 29.46747425202564,-81.2442545563 29.472433243399976,-81.2815606519 29.47653024280001,-81.2752052098 29.496236469699994,-81.2409002435 29.4828225312,-81.22571922956847 29.479108663162705,-81.2273819469 29.49025524130002,-81.2376861542 29.52173806820001,-81.21436262759998 29.524373685799986,-81.21508686530001 29.491644609199994,-81.21274535382314 29.475934747199815,-81.20251445653352 29.47343187073794,-81.1778413341 29.488827838299983,-81.15084304890001 29.511618865999992,-81.13714804159999 29.494980623199993,-81.17061893429998 29.480056431500003))"
	//wktStr := "POLYGON ((-72.6625489755 41.7058281537, -72.65773657440289 41.72875037718501, -72.67968755360832 41.7313294479, -72.6857425043 41.7313294479, -72.7291221889 41.7264827807, -72.72912675064119 41.73325756824532, -72.7382868702 41.7333095453, -72.7340325441 41.7535300403, -72.6921720683 41.7437211, -72.67855597165304 41.7421213071, -72.65492941750408 41.7421213071, -72.648843517 41.7711094086, -72.6485312306 41.8038437256, -72.6214112834 41.8006735846, -72.6345482685 41.7694384099, -72.6402894289542 41.7421213071, -72.6131485194 41.7421213071, -72.5697550285 41.7469517602, -72.56976368878595 41.734090092905774, -72.5685049073 41.734082563, -72.56976771291183 41.72811373503412, -72.5697688111 41.7264827807, -72.5701048260383 41.72652032252621, -72.5727826404 41.7138632571, -72.6146185521 41.7236843422, -72.64345228301231 41.72707208314331, -72.648268316 41.7041568434, -72.6485812398 41.6714225324, -72.6756614172 41.6745917717, -72.6625489755 41.7058281537))"
	//wktStr := "POLYGON ((2.36258301547 49.1554028788, 2.32450743259275 49.159886110452945, 2.28594277525 49.1393939631, 2.25221630224 49.1152173791, 2.2324758142 49.1311002013, 2.27553945935 49.1477694931, 2.3030857285023614 49.16240842146123, 2.26947440243 49.1663660066, 2.21943170346 49.1673335824, 2.219582492520187 49.16788276114527, 2.21934076418 49.1678849444, 2.22438624324 49.1880852857, 2.2720718897 49.1780824974, 2.322427514445812 49.172687236978426, 2.34992331879 49.1872993468, 2.38371035179 49.2114691083, 2.40345296463 49.1955818574, 2.36033669372 49.1789247471, 2.3442067203687755 49.17035374417951, 2.36471025897 49.1681569304, 2.41477077971 49.1676879929, 2.414112002116976 49.16506821604261, 2.41554708482 49.1650399948, 2.40998264253 49.1448958544, 2.36258301547 49.1554028788))"
	//wktStr := "POLYGON ((0.2711644403075005 51.54312059842762, 0.2563670864183439 51.52901381380369, 0.279320203613 51.5112520394, 0.317957540861 51.4890466335, 0.292325658465 51.4762236511, 0.265807604962 51.5044896103, 0.2463956227059696 51.519507702641555, 0.225465131514 51.4995540046, 0.2188174638507491 51.490519604605005, 0.212066588995 51.4708284321, 0.2054624744272602 51.47236973272622, 0.203567394948 51.4697942571, 0.1860944282249177 51.47688994227653, 0.181282434315 51.478012989, 0.1818266024812378 51.47862308515634, 0.176029779341 51.4809771463, 0.1970984368627807 51.495745158872936, 0.206808199108 51.5066312948, 0.2124320132035638 51.51596418009461, 0.199683884274 51.5100697613, 0.162590405921 51.4868674146, 0.142958395423 51.5032922814, 0.189339159419 51.5187311785, 0.2236638404136843 51.53460372849, 0.2246873700220026 51.53630230605102, 0.207185647425 51.5498425248, 0.168479893897 51.5720366198, 0.194133587039 51.5848653939, 0.220712347981 51.5566044637, 0.2317665475698583 51.54805040990218, 0.23900050938 51.5600553828, 0.250004068882 51.5920682784, 0.280838285059 51.5848796543, 0.255253628416 51.5562665204, 0.2493554739328017 51.54648415171724, 0.2583223221833852 51.55063063590837, 0.264415919868 51.5564416226, 0.28636573264 51.5861984178, 0.313937245677 51.5750106551, 0.3071167120774645 51.57024091211699, 0.319848489039 51.5595819006, 0.273411733954 51.5441596913, 0.2711644403075005 51.54312059842762))"
	//wktStr := "POLYGON ((0.2008572048150811 51.709533438777214, 0.200896247719 51.7095814696, 0.2132117574376404 51.704186333748105, 0.22701657809 51.6982115463, 0.2269500024205612 51.69816793095822, 0.227858788727 51.697769813, 0.2235238137622581 51.69496201358428, 0.223692836799 51.6948891397, 0.1890749493792964 51.67212499991679, 0.1747584155074132 51.659551135620525, 0.202557835531 51.6572078248, 0.255321187152 51.6576331985, 0.250858061012 51.6373513756, 0.20021683109 51.6465141528, 0.1634438872318165 51.64961385999315, 0.1626223260587653 51.648892302750774, 0.136474056847285 51.62555487711058, 0.113814799153 51.5976092631, 0.1135776205041143 51.59771322440015, 0.110556610886 51.5939163574, 0.1099715567535497 51.594168754122876, 0.109968052569 51.5941643708, 0.1065640249594312 51.595638788791874, 0.0834824667829 51.6055963454, 0.0835190629422136 51.60562046653286, 0.0829291811372 51.6058759675, 0.0877432758871456 51.60903706011458, 0.0868804941644 51.6094152388, 0.1220918134233861 51.63227504187265, 0.1434835719255599 51.6512963778022, 0.106912490209 51.6543790694, 0.0541533775094 51.6539332449, 0.0585745794774 51.6742166437, 0.109255580576 51.6650725998, 0.1546712552189607 51.661244352961674, 0.1744837707688314 51.6788614419843, 0.196589443265 51.706574704, 0.197086725827796 51.706360301794035, 0.199945799337 51.7099278991, 0.2008572048150811 51.709533438777214))"
	//wktStr := "POLYGON ((0.11166115001 51.7414397683, 0.1472144433465579 51.727567070234024, 0.160070343004 51.7562385744, 0.166544800334 51.7887276383, 0.1677107619660893 51.78852709530657, 0.167991525295 51.7899182524, 0.199846852447 51.7844065439, 0.178226339468 51.7545304245, 0.1632319209782327 51.721317139211905, 0.195866235259 51.7085834106, 0.244224363134 51.6954157559, 0.226610240871 51.6780981336, 0.186586785173 51.6994513955, 0.158336518085603 51.7104736101358, 0.1501839309509339 51.69241527712845, 0.144251175206 51.6625704884, 0.1426411368027305 51.66284778104926, 0.140774867104 51.6535802441, 0.108985954794 51.6590903703, 0.1151191650144336 51.6675878173456, 0.112441734562 51.6680489441, 0.133934160839 51.6979490983, 0.1423472054636274 51.71671205153011, 0.102372831586 51.7323085812, 0.0539642117702 51.745460682, 0.0715676306257 51.762782983, 0.11166115001 51.7414397683))"
	//wktStr := "POLYGON ((-0.833788118399 52.2902366434, -0.8026249796093073 52.30537825953708, -0.778218593848 52.3379753068, -0.763349861813 52.3694254936, -0.732886261322 52.3609017084, -0.7357656335184846 52.35820477863703, -0.733767102567 52.3576475343, -0.762978224636 52.3302044668, -0.7694921520675876 52.32147691248595, -0.757704308286 52.3272044166, -0.721135765466 52.3511290064, -0.700311342297 52.3351002406, -0.74671948696 52.3187553903, -0.7856337495178369 52.29984991864204, -0.803723177678 52.2756131631, -0.818483560106 52.2441494272, -0.820447198455545 52.24469754848002, -0.821308625044 52.2428708601, -0.851719449911 52.2513912319, -0.82251326472 52.278815565, -0.8188447359102825 52.28371523380657, -0.822813799876 52.2817869684, -0.859309163387 52.2578532674, -0.880134718091 52.2738763709, -0.833788118399 52.2902366434))"
	//wktStr := "POLYGON ((-0.982005150979 52.5991320412, -1.016604521183245 52.60140167563649, -0.99941261584 52.5838882996, -0.965303167132 52.5585541112, -0.9665650411364026 52.55809700775142, -0.966288046815 52.5578883901, -0.995335183383 52.5475283886, -1.01542819587 52.5778993993, -1.0313692449574456 52.59438714979867, -1.05137759147 52.5794703999, -1.07883458775 52.5513094808, -1.0921548091990196 52.55789956208104, -1.0924054629 52.557715138, -1.0941065559864065 52.558865174403486, -1.10495658176 52.5642331446, -1.1033572212922322 52.5651191542375, -1.11492588586 52.5729402344, -1.06997309759 52.5909821025, -1.0471935246466306 52.60340823977995, -1.08406902177 52.6058271826, -1.13792603611 52.6044845227, -1.13432709657 52.6248356681, -1.08215863362 52.6165565145, -1.050815662483109 52.614500486035034, -1.06418957296 52.6283330574, -1.0878009521279728 52.64607218680709, -1.0976008209 52.6533329819, -1.0975093728113683 52.65336608224692, -1.09808591133 52.6537992332, -1.0689996626 52.6641645584, -1.0687433585587447 52.66377816483253, -1.06863206832 52.6638184472, -1.04814729744 52.6335342689, -1.0285548983796777 52.613575512418414, -1.0285463338555432 52.6135801843213, -1.00448415046 52.6315233848, -0.976962067244 52.6596789961, -0.9658803866268392 52.654198617471025, -0.965650379406 52.6543674293, -0.9640951115537727 52.65331572040015, -0.950817445388 52.6467493303, -0.9527804161614125 52.645664455811705, -0.943126057968 52.6391359505, -0.988145254705 52.6211096656, -1.0056902791774422 52.61154036230701, -0.980098416017 52.6098615936, -0.926233577358 52.6111860477, -0.92987457857 52.5908360119, -0.982005150979 52.5991320412))"
	//wktStr := "POLYGON ((-2.41046130521 53.4447638594, -2.40115671751158 53.465252145434725, -2.43826444769 53.4624829177, -2.49120175074 53.4536447092, -2.49550281577 53.4739539358, -2.494336273624496 53.47393775008199, -2.4946064044 53.4754268159, -2.43965226373 53.4740802127, -2.4058442960691866 53.476239089616726, -2.42715465189 53.4892441924, -2.4709880342 53.5090370186, -2.44636171829 53.5233313911, -2.41416752044 53.4967786654, -2.3927704948566673 53.48371823271057, -2.38471691787 53.5014518491, -2.37835584523 53.5339677717, -2.34512728899 53.5286225884, -2.36720299443 53.4986346916, -2.376532182107302 53.478110874600894, -2.3630940036351693 53.478968996970075, -2.34119010363 53.4806036194, -2.28821010726 53.4894225622, -2.28395507689 53.4691118481, -2.2849647835052513 53.46912624367569, -2.28480565312 53.468237615, -2.33975073599 53.4696058325, -2.360998868575493 53.46824899349054, -2.366726794209757 53.4678215372583, -2.34661968626 53.4555484515, -2.30285995625 53.4357411706, -2.32747378167 53.4214533025, -2.3595934507 53.4480135169, -2.383563684613912 53.462641867473344, -2.39297089833 53.4419463965, -2.39931493126 53.4094302479, -2.4324793448 53.4147722896, -2.41046130521 53.4447638594))"
	//wktStr := "POLYGON ((-5.21374146374 50.0951404316, -5.238830050611195 50.08708750920928, -5.20040784843 50.076627848, -5.20697244211 50.0666924599, -5.242843466371891 50.07645725064966, -5.24217248164 50.0601727584, -5.242192797374599 50.06017241212519, -5.24193381513 50.0537698278, -5.25873601561 50.0534885856, -5.2598295651371805 50.08029101449301, -5.28874151468 50.0708222843, -5.288896531489044 50.07101719456958, -5.29705521005 50.0683984259, -5.30457743751 50.078050178, -5.278933687805513 50.08628171113442, -5.30790280908 50.0941676696, -5.30132957716 50.1041017311, -5.260350497665205 50.09294601217056, -5.26133869528 50.1167266337, -5.24451452336 50.1170130387, -5.244394924977679 50.11411044071616, -5.24437458508 50.1141107807, -5.2437102428110425 50.09768684629331, -5.21690261673 50.1064668645, -5.20925278103 50.096855028, -5.213893397178323 50.09533521550266, -5.21374146374 50.0951404316))"
	//wktStr := "POLYGON ((-5.28570108036 50.0044916169, -5.255622999087868 50.0044916169, -5.27516422812 50.0140855066, -5.26496828443 50.0226613107, -5.239276707614113 50.0100465561579, -5.23985598844 50.0336917018, -5.239739324961099 50.03369287330468, -5.23975980004 50.0347077636, -5.22295987859 50.0348464671, -5.2225804746606865 50.01570914695762, -5.19944779262 50.0308055358, -5.18747860042 50.023233315, -5.2162017704076895 50.0044916169, -5.18369891964 50.0044916169, -5.18369689273 49.9936997577, -5.205984341371605 49.9936997577, -5.18828246772 49.9850080066, -5.19846864966 49.9764313284, -5.221948574365192 49.98795894559416, -5.22148607709 49.96880771, -5.23826226622 49.9686389807, -5.238534569550164 49.979753886653796, -5.23865110416 49.9797529234, -5.238852034422616 49.98971245393731, -5.26386717215 49.9733902249, -5.2758212105 49.9809642415, -5.2563061695220155 49.9936997577, -5.28570310727 49.9936997577, -5.28570108036 50.0044916169))"
	//wktStr := "POLYGON ((-5.6802754132 50.1331644392, -5.672416072288584 50.120472524859515, -5.66575045857 50.1324186176, -5.65582801126 50.1645303234, -5.62575868984 50.1576295173, -5.64990049557 50.1287814027, -5.658505576067526 50.11336853820408, -5.635106926 50.1202740772, -5.59237552852 50.1381960308, -5.57904740257 50.1195989442, -5.62807148183 50.1104704461, -5.654136234977486 50.102778405284624, -5.62709093409 50.0953983303, -5.57770629082 50.0871676545, -5.59021112203 50.0683342709, -5.63367506213 50.0854675645, -5.654209285050764 50.09107069757623, -5.64507388562 50.0763180983, -5.61945475939 50.0480142609, -5.64908846029 50.0404583956, -5.66070074407 50.0723332392, -5.671596964878737 50.08992008796952, -5.68066430918 50.0736792464, -5.69056292501 50.0415669012, -5.72058416215 50.0484641756, -5.69649586802 50.0773167878, -5.687838539536169 50.09283242571709, -5.71330854982 50.0853159011, -5.75598011807 50.067381518, -5.76932985711 50.0859746468, -5.72033849788 50.0951200644, -5.695698814542276 50.10239186399719, -5.71731766515 50.1082909573, -5.76673149077 50.1165033175, -5.75425165777 50.1353407446, -5.7107291076 50.1182213706, -5.6847371225806205 50.11112872262983, -5.69592117288 50.1291801374, -5.72159645581 50.1574792396, -5.69191517165 50.1650385364, -5.6802754132 50.1331644392))"
	wktStr := "POLYGON( (8.62883106519 50.0445149656, 8.534466861444404 50.02180734290968, 8.53497888571 49.9715565954, 8.51820004288 49.9714847, 8.517712187759079 50.01777553006917, 8.49779800804 50.0129834197, 8.49191126524 50.0230906883, 8.494910001894636 50.023812319904565, 8.49233921324 50.0282606405, 8.51753828493684 50.034276547784295, 8.517518586614017 50.036145651285224, 8.46067995035 50.0224946542, 8.45479960543 50.0326037399, 8.51739746573664 50.04763837876119, 8.51725514352 50.0611428256, 8.534065325 50.0612146676, 8.534162630197835 50.05166500510094, 8.570129366 50.0603034295, 8.57602082208 50.0501961968, 8.534279745046412 50.04017119693419, 8.53429903627058 50.0382779303452, 8.62344132599 50.0595593393, 8.62930301272 50.0494449958, 8.62636176421957 50.0487428364738, 8.62883106519 50.0445149656))"
	var polygon = wkt2.MustPolygon(wkt2.WktReader{}.Read(wktStr))
	ps, err := polygon.Split(21)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(len(ps))
	//ps = append(ps, *polygon)
	clipboard.WriteAll(geojson.GeojsonWriter{}.EncodePolygons(ps))
	total := polygon.GetArea() * geo.EarthRadius2
	fmt.Println("total_area=", total, ", total_point_cnt=", len(polygon.GetShell()))
	for _, p := range ps {
		area := p.GetArea() * geo.EarthRadius2
		fmt.Println(math.Round(area/total*10000)/100, "%", len(p.GetShell()), area)
	}
}
