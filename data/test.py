

from shapely.wkb import loads
import binascii

# Your WKB data as a hex string
wkb_hex = "30313033303030303230453631303030303030313030303030303237303030303030313846333533374338444245354634304545444334443836393243383432343046324339323034323933424535463430323733333634433639324338343234303246343944353846394442453546343037363039464635303933433834323430433531443631423739444245354634304446454238423535393343383432343039374546314438413946424535463430454532303934323139334338343234303244464345343332413042453546343034434433424630453933433834323430374637423746383741304245354634303946393132374645393243383432343036393833303842434134424535463430313546393446323339324338343234304635323238413239413842453546343035443232443339373931433834323430423231414342413141384245354634303238324142463632393043383432343039464631364545324138424535463430353833443245423438464338343234303644463434434631413842453546343042313946344438323846433834323430384636384646454541384245354634303437353239424536384543383432343043454243413546304138424535463430314344443542424438434338343234303434313337304633413842453546343042363031364234463839433834323430313635384339463341384245354634303144423034374532383843383432343038413934373746394138424535463430374641463938453538314338343234303530424639393943413742453546343035303637454346373746433834323430353637364444413541374245354634304436383842353533374443383432343033333939333941394137424535463430354231333434363037434338343234304638423430463741413742453546343044344437354142313742433834323430454535303933464441364245354634303331363736344134374143383432343043433037393941374136424535463430463536314136413937414338343234303431364244314630413542453546343046434244454142343741433834323430434244453343344241344245354634303130324343314345374143383432343042323338433743424132424535463430374630333742383437384338343234304639304134304138413042453546343031353032353238453738433834323430423735434342363241304245354634304536303339343846373843383432343036303643373834413945424535463430434242433331393937384338343234303530353532414345394242453546343041334145374139443738433834323430334243364141363539424245354634304334313745383934373743383432343030444334453230323941424535463430363736384133413437374338343234303237343838434344393442453546343041414330443044463737433834323430363936333134444639334245354634304335454135384630373743383432343046323445424533313930424535463430343942453441314137384338343234303138413142334445384642453546343041414131363834453738433834323430353741453242433038444245354634303646453445314136374343383432343038394544444539393844424535463430383337464232303038394338343234303138463335333743384442453546343045454443344438363932433834323430"

# Convert the hex string to binary
wkb_bin = binascii.unhexlify(wkb_hex)

# Attempt to load the geometry
try:
    geom = loads(wkb_bin)
    print(f"Geometry loaded successfully: {geom}")
except Exception as e:
    print(f"Error loading geometry: {e}")