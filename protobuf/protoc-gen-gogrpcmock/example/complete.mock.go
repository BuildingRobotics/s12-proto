// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: complete.proto

package s12_complete

import context "context"
import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/SafetyCulture/s12-proto/protobuf/s12proto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type MyCompleteServiceMock struct{}

func (m *MyCompleteServiceMock) GetMessage(ctx context.Context, req *MsgRequest) (*MsgResponse, error) {
	res :=
		&MsgResponse{
			Id: "c902b51f-4243-45f4-b274-5d66e5844920",
			Completed: []*Complete{
				&Complete{
					Id:          "5c1c92b0-2122-4a39-af75-8cb53242d420",
					Status:      "in progress",
					Description: "facere repellendus ex nisi ut nobis. doloremque aut autem unde labore exercitationem. voluptates et dolores nemo ut. veritatis voluptatem autem voluptate rerum nihil qui ut! et commodi quasi ab consequatur dolores sit! tempore tempore quibusdam impedit odit ad excepturi voluptatem. temporibus ex saepe magnam atque dolorem excepturi. amet quas nobis odit dicta praesentium. ipsa nisi similique qui deserunt. aperiam harum sit.",
					Email:       "BenjaminRamirez@Quaxo.biz",
					NotEmail:    "qui",
					Phone:       "206-10-52",
					Word: []string{
						"id",
						"laboriosam",
						"est",
						"officia",
						"aliquam",
						"sit",
						"ducimus",
						"repellat",
					},
					Url:          "https://quire.gov/neque/beatae/qui",
					SingleNumber: 102,
					RepeatedNumber: []int64{
						457,
						428,
						153,
					},
					Lat:       -81,
					Lng:       50,
					Words:     "nobis aspernatur",
					Wordsn:    "eos illum ducimus est nobis accusantium qui similique fugit sint",
					Intn:      0,
					Fullname:  "Sarah Hawkins",
					Firstname: "Lisa",
					Lastname:  "Gilbert",
					Paragraph: "error quos distinctio quidem nesciunt blanditiis! omnis ut expedita fuga totam nihil adipisci ex consectetur occaecati. quo libero eius sint animi id. id provident vel fuga cupiditate ipsam nemo cupiditate veniam. quasi modi consequuntur facilis expedita fuga. tempore quia molestiae aut.",
					Paragraphs: "ut voluptate dignissimos atque molestias. laudantium pariatur dolore quia aut natus cupiditate. corrupti ratione quam accusamus vel. et adipisci necessitatibus error ducimus et error iusto minus. fuga ut error et. tempore perferendis illo. ratione quia id.	error iure est illum cupiditate nesciunt eaque. provident veritatis voluptate voluptates velit non. consequuntur fugiat quaerat velit delectus est voluptas! optio consequatur veniam impedit et voluptatem rem id. distinctio consequatur aut assumenda qui neque. voluptatum nemo non aut dicta. labore magnam consectetur facere esse. placeat numquam non. ratione laborum maiores culpa ut.	debitis ullam et velit itaque rerum! amet recusandae est iure fuga molestiae. eos nihil repudiandae atque illo dignissimos. deserunt illo dicta quia! et recusandae autem consequuntur consectetur quaerat. quia est excepturi. dicta minima vel velit. et non tempore quidem quaerat. alias ea fugit rerum et sapiente. non autem voluptatem nobis nostrum eius.	in modi consequuntur!",
					Paragraphsn: "nesciunt accusantium itaque similique.	minima architecto et consequuntur perspiciatis aliquam veniam. vitae alias omnis neque porro consequatur. doloribus voluptatem commodi distinctio accusantium. et commodi eos velit quod. praesentium exercitationem quibusdam sit odit ea praesentium. et nulla excepturi aut aut et sed qui. incidunt dolores et. fugit vel accusantium ex.",
					Uuid:         "2371b544-36f3-44d2-9910-d118c77860b4",
					EmailAddress: "ShawnPierce@Ntag.name",
					PhoneNumber:  "8-324-724-85-19",
					Company:      "Blognation",
					Brand:        "Skimia",
					Product:      "Performance Digital Compressor",
					Color:        "Crimson",
					Hexcolor:     "#8c809b",
					Latitude:     72.971542,
					Longitude:    -76.059044,
					Floatn:       0.3150,
					Boolean:      true,
					SomeOneof: &Complete_OrCouldBe{
						OrCouldBe: &SubMessage{
							Id: "1f46ee8c-381e-4ec2-84d6-ab4fd193b183",
						},
					},
					AnotherOneof: &Complete_This{
						This: "id",
					},
				},
				&Complete{
					Id:          "ec4a7c07-81b9-41de-8fbd-3fecce55e1da",
					Status:      "in progress",
					Description: "blanditiis nihil consequatur voluptate ipsum molestiae.",
					Email:       "CharlesThompson@Skynoodle.com",
					NotEmail:    "a",
					Phone:       "285-94-89",
					Word: []string{
						"quam",
						"necessitatibus",
						"praesentium",
						"autem",
						"voluptatem",
						"dolorem",
						"error",
						"neque",
					},
					Url:          "https://divanoodle.com/voluptatibus",
					SingleNumber: 368,
					RepeatedNumber: []int64{
						217,
						718,
						814,
					},
					Lat:       -22,
					Lng:       62,
					Words:     "at ipsa veritatis excepturi",
					Wordsn:    "reprehenderit quibusdam non laborum autem consequuntur molestiae et vel eaque",
					Intn:      0,
					Fullname:  "Mr. Dr. Keith Adams",
					Firstname: "Lillian",
					Lastname:  "Gonzales",
					Paragraph: "reiciendis autem voluptas dolores tempore quis. minima in quia qui earum. impedit voluptatum ab ex quasi doloribus qui. numquam mollitia ut qui sit. debitis maiores fugit veniam pariatur non ad. quam recusandae vitae quia eos in. voluptas fugit vitae repudiandae illum sequi dolores voluptatem ipsum libero! quibusdam inventore sint unde et cupiditate ad velit. porro sed sed ipsum corporis voluptatibus. ex cupiditate ea reiciendis similique ipsam.",
					Paragraphs: "ut qui ullam inventore.	praesentium repellat aut. voluptatem eos sit laudantium consequatur sed totam autem est. sed recusandae sit ipsum ut voluptates. atque quam qui facere non ut aut dolores sunt temporibus. amet commodi accusantium quod incidunt beatae dolores a. hic quo totam in voluptate inventore sit asperiores esse quidem.	ut iste laboriosam assumenda unde nam tempore. deserunt consequuntur possimus ab amet. tenetur eum ut voluptatibus. iure culpa repudiandae voluptas est vitae! tempora id qui et rerum magnam. qui quasi tempore rerum perferendis velit laborum. quis eligendi quos eos unde ad. reprehenderit magni sed quaerat fugiat mollitia aut nisi.",
					Paragraphsn: "aut nisi consequuntur accusamus! esse rerum enim nobis consectetur aut tempore quaerat ea. quaerat eos illo ipsum earum perspiciatis. dolore et et et cum mollitia odit. in voluptas laborum. magnam itaque tenetur eligendi animi veniam odio sit. mollitia et ut adipisci quaerat qui sequi dolore fuga blanditiis non.	quia dolor aliquid aut reprehenderit illo aliquid qui voluptatem. tempore delectus iusto consequatur qui ut. sed et ut. ipsum quo amet perspiciatis et sed. voluptatum nobis optio odio. cum explicabo magni maiores. aut earum rerum. illum consectetur dolor sint. dolor veniam voluptate voluptatem fugit molestiae officia repellat. molestias architecto vero est eos vero quas.",
					Uuid:         "c90c5b6a-17eb-45a3-b4ec-4770af09fa8e",
					EmailAddress: "5Cunningham@Photobug.org",
					PhoneNumber:  "4-050-399-34-72",
					Company:      "Roodel",
					Brand:        "Voonder",
					Product:      "Audible Disc Filter",
					Color:        "Turquoise",
					Hexcolor:     "#dc6a64",
					Latitude:     72.218689,
					Longitude:    178.292694,
					Floatn:       1.9063,
					Boolean:      true,
					SomeOneof: &Complete_OrCouldBe{
						OrCouldBe: &SubMessage{
							Id: "e4d16ca6-54cd-4b06-b1ce-2189d3af4c98",
						},
					},
					AnotherOneof: &Complete_This{
						This: "vero",
					},
				},
				&Complete{
					Id:          "916cad21-8f3d-4c9f-9c4a-b274cfbc020c",
					Status:      "in progress",
					Description: "velit eum consequatur qui et. amet nesciunt quisquam harum cum et. sit asperiores optio illo ipsam itaque nulla. vel est nihil! reiciendis fuga rerum a. velit exercitationem repellat sunt placeat ipsam. sint minus deserunt illo. voluptatum fuga sint.",
					Email:       "NormaCarr@Tazzy.name",
					NotEmail:    "quasi",
					Phone:       "415-32-88",
					Word: []string{
						"dignissimos",
						"consequatur",
						"sunt",
						"ratione",
						"vero",
						"non",
						"veritatis",
						"blanditiis",
					},
					Url:          "https://twitterwire.net/eveniet",
					SingleNumber: 325,
					RepeatedNumber: []int64{
						849,
						473,
						722,
					},
					Lat:       55,
					Lng:       -14,
					Words:     "aut eligendi distinctio",
					Wordsn:    "qui totam reprehenderit qui officia error ipsum magnam ducimus consequatur",
					Intn:      0,
					Fullname:  "Samuel Cole",
					Firstname: "Robin",
					Lastname:  "Henderson",
					Paragraph: "vel est eos deleniti vero reprehenderit occaecati. eum ad quas aspernatur quod.",
					Paragraphs: "quis illum eos distinctio quis quidem quod. iusto dolor maxime saepe voluptatum placeat. et in voluptatem nam. sit quo sit quia et.	debitis dolores blanditiis ratione hic corrupti quia quis iusto. fugiat laborum maiores est aspernatur doloribus. omnis odio est molestiae distinctio ut molestias. omnis veniam repellendus et atque. enim laudantium aliquid earum voluptatem! temporibus officiis vel ipsam in ipsa est. ducimus voluptates illum iste sunt porro nobis harum. commodi eum non amet error odit ipsa.	numquam est officia labore. vero culpa consequatur quas aut nam. qui quis hic. optio numquam et repellat soluta repudiandae inventore rerum modi! qui ut et voluptatem! sapiente et qui rerum quidem sunt in molestias nesciunt sit. minus ex velit sed provident in beatae. quidem ducimus voluptatum voluptatem sint consequatur suscipit. quae illo voluptatem eos vitae. voluptas illum enim enim.",
					Paragraphsn: "laborum at voluptas placeat libero aut vitae. facilis omnis dolor nemo sapiente consequuntur voluptas quaerat corrupti. asperiores mollitia est assumenda odio ullam. explicabo in sapiente qui molestias alias nam.	voluptatem blanditiis consequatur sed quo. id velit cum expedita ducimus quasi! doloribus unde facere tempore sunt. et aspernatur odio fugit consequatur ex. harum impedit quas earum optio est quo cum. animi animi repellendus dolor et voluptate. saepe velit quaerat molestias ipsam. consectetur ex pariatur nostrum consequatur tenetur ducimus.",
					Uuid:         "b1ff0086-94fa-4a8f-9995-236dc80e9ca5",
					EmailAddress: "quo@Zoomcast.net",
					PhoneNumber:  "244-54-45",
					Company:      "Flipbug",
					Brand:        "Demivee",
					Product:      "Audible Bridge",
					Color:        "Crimson",
					Hexcolor:     "#5c8e64",
					Latitude:     46.638931,
					Longitude:    -11.375198,
					Floatn:       0.1726,
					Boolean:      true,
					SomeOneof: &Complete_OrCouldBe{
						OrCouldBe: &SubMessage{
							Id: "3dba3fd6-e21d-48db-b02e-3069e5bac14a",
						},
					},
					AnotherOneof: &Complete_This{
						This: "voluptas",
					},
				},
				&Complete{
					Id:          "dd9e2c73-4324-419d-bc97-ebecf3d17fcf",
					Status:      "complete",
					Description: "amet necessitatibus deleniti nulla voluptate. tempore fugiat ratione beatae optio dolor cumque. id consequatur minus molestiae dolores eum. et qui rerum distinctio velit facere quo accusamus quaerat. ad sit est quo vel cum. expedita placeat et consequuntur.",
					Email:       "qui@Brainlounge.gov",
					NotEmail:    "provident",
					Phone:       "7-546-489-05-82",
					Word: []string{
						"praesentium",
						"et",
						"facere",
						"nobis",
						"minima",
						"tenetur",
						"consequatur",
						"consequatur",
					},
					Url:          "https://rhyloo.com/vero/voluptas",
					SingleNumber: 842,
					RepeatedNumber: []int64{
						884,
						250,
						933,
					},
					Lat:       62,
					Lng:       42,
					Words:     "dignissimos et iste magnam",
					Wordsn:    "ratione suscipit itaque qui optio voluptatibus voluptatum ut unde non",
					Intn:      2,
					Fullname:  "Mr. Dr. Mark Martinez",
					Firstname: "Douglas",
					Lastname:  "Myers",
					Paragraph: "ad nihil est sapiente ea ullam odio excepturi quam. alias id adipisci cumque asperiores corrupti explicabo. fuga quo molestiae sequi voluptas laudantium autem vero. minima ipsa earum assumenda magni enim dolores. non provident fuga quia. consequuntur sequi in nihil accusantium. quaerat distinctio dolor fuga nesciunt. numquam saepe et sed. dolorum accusantium veritatis earum blanditiis. dolorem doloremque eum facere provident.",
					Paragraphs: "quos explicabo harum omnis assumenda. quos libero repellendus porro incidunt molestiae tenetur. quia accusamus aperiam. dolores quae qui ab repellendus in ut eligendi. deserunt corrupti consequatur rerum est cumque. qui error non magni aut blanditiis at commodi incidunt quos. quis et nostrum eligendi ipsam ipsum! totam qui qui commodi quasi enim quos.	quisquam enim beatae fugiat! praesentium reiciendis sint non voluptas at. laudantium voluptas dolorem illum et expedita. fugit distinctio non voluptatem amet atque. dolor praesentium consequatur ut quis in soluta perferendis. occaecati beatae nulla magni quas enim molestiae.",
					Paragraphsn: "autem sequi voluptate veritatis eligendi dolorum. qui rerum aut harum ad voluptates. veritatis quia tenetur asperiores error reprehenderit quas. atque tempore deserunt dolorum minus. qui sapiente velit eos repudiandae aspernatur saepe.	et saepe velit doloremque ipsa nesciunt sit delectus saepe. reiciendis eos velit nulla ut consequatur ex neque consectetur. ut illum facilis placeat mollitia praesentium. vitae voluptatem culpa dolores. dicta sapiente minus id accusantium consequatur necessitatibus. voluptatibus aut harum officiis dolorem et voluptatem dicta. aliquid voluptas corrupti mollitia eius cumque sed quam quis esse. amet ex similique tenetur non non quas voluptatem. deleniti aut repellat accusamus laborum numquam!",
					Uuid:         "fcc7cf7c-92eb-434b-bdb1-52a1da41f6ee",
					EmailAddress: "RandyRivera@Realcube.gov",
					PhoneNumber:  "591-03-17",
					Company:      "Feedfire",
					Brand:        "Rhybox",
					Product:      "Auto Tuner",
					Color:        "Aquamarine",
					Hexcolor:     "#3d62e4",
					Latitude:     -25.718605,
					Longitude:    -42.073151,
					Floatn:       2.8207,
					Boolean:      false,
					SomeOneof: &Complete_OrCouldBe{
						OrCouldBe: &SubMessage{
							Id: "333710e4-946e-4cea-bf84-2c9aac13c94d",
						},
					},
					AnotherOneof: &Complete_This{
						This: "autem",
					},
				},
				&Complete{
					Id:          "f70f6955-e791-4a4a-a1f8-7671e52df7b4",
					Status:      "complete",
					Description: "est saepe facilis nam necessitatibus harum. nihil optio quo id deleniti! sunt quibusdam accusantium. qui officiis cum iste. et molestias velit voluptas non et eos optio tenetur. ea inventore fuga optio voluptate dolore odio. aspernatur ex similique repudiandae architecto. commodi dignissimos velit reprehenderit eum ab consequatur.",
					Email:       "8Williams@Topicware.info",
					NotEmail:    "nisi",
					Phone:       "793-28-05",
					Word: []string{
						"vel",
						"est",
						"repudiandae",
						"commodi",
						"magni",
						"sed",
						"quia",
						"aperiam",
					},
					Url:          "https://agimba.edu/esse/est/laudantium/et",
					SingleNumber: 859,
					RepeatedNumber: []int64{
						350,
						272,
						774,
					},
					Lat:       25,
					Lng:       0,
					Words:     "totam quas expedita",
					Wordsn:    "recusandae vel dolore modi dolores corrupti deleniti qui quia beatae",
					Intn:      4,
					Fullname:  "William Payne",
					Firstname: "Jimmy",
					Lastname:  "Dixon",
					Paragraph: "facere voluptas quisquam cupiditate. et voluptas quo sint beatae id non quidem dolorem ducimus nostrum. adipisci aut qui quas ea.",
					Paragraphs: "ipsam ad velit sit non. eum eos quibusdam non quam. itaque qui veritatis porro laborum. aut vel nihil molestias architecto similique itaque. exercitationem necessitatibus provident et saepe sit sunt. corrupti quaerat molestias consequuntur dolor est veritatis.	hic quo alias adipisci numquam similique et magni et. iste eius odit rerum cumque minima. quas reiciendis et magni sequi non. est ut voluptatem inventore est non fuga! dolores maiores quis laudantium accusantium ut libero. dolore facilis facilis non est. sint amet assumenda sed voluptates a et.	qui vel at impedit incidunt tenetur exercitationem. quis vel cum quam pariatur enim. ut temporibus voluptatem rerum sint ut est ea quia. repudiandae vel facere eos et velit. facere commodi molestiae quasi perspiciatis ut. rerum voluptatem repudiandae architecto. earum voluptas pariatur eaque omnis beatae sed. atque sequi laudantium voluptatem ab quasi doloremque accusamus. aut aspernatur beatae est. vitae sint rerum deleniti non est doloremque.	nostrum praesentium consequatur. expedita quos sed assumenda libero quas facere. dolores et aut aperiam recusandae omnis repudiandae quia! aut quia quaerat consequatur. ex laudantium alias. eos mollitia et earum quia eius et.	molestias quis repudiandae eligendi. nihil sequi vero consequatur non est. est magni eum voluptatem. aspernatur delectus commodi! amet occaecati quo qui voluptate maiores similique placeat. autem voluptas a eveniet commodi et neque.",
					Paragraphsn: "possimus dignissimos qui quia reprehenderit eos nulla. nam quos odit asperiores. et ab molestiae. non aut facilis veritatis cupiditate. asperiores ea laboriosam voluptas dolore pariatur. placeat atque consequatur sapiente sed. repudiandae culpa libero occaecati exercitationem et. rem sed perferendis accusamus id. velit ab veritatis voluptas. nam minima voluptas error voluptatum vitae repellendus magni.	sint eum quidem corporis iste minima. eveniet voluptatem vel nam laborum. aut saepe dolorem vitae delectus. ab minima minima accusamus aut totam velit. eaque odio mollitia hic accusantium aliquam ullam.",
					Uuid:         "9fce4ecb-c4b7-47da-9096-896877d8bdac",
					EmailAddress: "minima_dolor@Yadel.mil",
					PhoneNumber:  "6-209-946-06-77",
					Company:      "Vinte",
					Brand:        "Twitterworks",
					Product:      "Performance Performance Viewer",
					Color:        "Aquamarine",
					Hexcolor:     "#df7902",
					Latitude:     -54.554470,
					Longitude:    92.423370,
					Floatn:       0.6251,
					Boolean:      true,
					SomeOneof: &Complete_CouldBe{
						CouldBe: "dolore",
					},
					AnotherOneof: &Complete_That{
						That: "ut earum et est deleniti. et quae distinctio laudantium et a. ex harum deleniti corporis! voluptates dolor sit nobis sit eaque illum.",
					},
				},
				&Complete{
					Id:          "f691090a-6366-4c72-bd6b-27216c475172",
					Status:      "in progress",
					Description: "perferendis voluptas dolorem architecto assumenda corrupti. iure eum tempora recusandae porro similique sed. dolor quo ducimus non magni. ut officiis voluptatibus iure omnis id et mollitia corrupti. sit eveniet accusamus et fugit. qui reprehenderit atque velit cumque libero eum! culpa consectetur temporibus exercitationem sapiente alias. voluptatem libero dolor atque omnis eaque. voluptatem cupiditate optio quia facere quod necessitatibus.",
					Email:       "qGarrett@Kazio.edu",
					NotEmail:    "itaque",
					Phone:       "2-800-798-16-66",
					Word: []string{
						"at",
						"architecto",
						"eius",
						"aut",
						"qui",
						"ipsam",
						"qui",
						"saepe",
					},
					Url:          "https://kwinu.edu/praesentium/enim/soluta/illo/eaque",
					SingleNumber: 345,
					RepeatedNumber: []int64{
						111,
						228,
						43,
					},
					Lat:       23,
					Lng:       118,
					Words:     "quos",
					Wordsn:    "repellat sed placeat debitis ratione facilis ex placeat cumque ut",
					Intn:      1,
					Fullname:  "Kevin Oliver",
					Firstname: "Beverly",
					Lastname:  "Dunn",
					Paragraph: "rem autem ut voluptatem occaecati harum architecto. vero dolores quia ducimus vitae voluptatem natus. ipsa harum voluptatem ipsa qui esse quasi nesciunt maxime recusandae iste. ipsum maiores sed dolores omnis commodi!",
					Paragraphs: "illum non ex necessitatibus et ad et blanditiis reprehenderit. molestiae voluptatum dignissimos enim aperiam esse et. repellat sapiente ea illo quam mollitia. voluptatibus esse dolor atque! sed nulla amet tempora nihil fuga ad. architecto dolorem numquam.	voluptatem quia maiores dolorum. aut suscipit quod fugiat et ut enim similique. dolores temporibus quas dignissimos fugit est.	suscipit rerum animi excepturi esse. eligendi modi veritatis quas dolor mollitia sint. et voluptas voluptatem. cumque nobis ab fugiat dolorum amet nihil! qui recusandae aspernatur. unde alias modi voluptatibus et sunt et culpa. beatae sit laudantium et. quidem cumque quae debitis in dolorum pariatur.	vero commodi similique ipsam et hic dignissimos doloremque. aliquam aut provident praesentium. facilis doloribus qui eos corrupti at.",
					Paragraphsn: "repudiandae exercitationem corporis expedita unde voluptatibus libero placeat nam. molestias nulla dolores fugiat. totam necessitatibus voluptatum asperiores quo labore qui deleniti modi non. enim ea explicabo officiis velit molestiae. ut eos alias.	veritatis accusantium necessitatibus perferendis. mollitia rem qui dolores quas aut. porro ducimus qui nemo repellat.",
					Uuid:         "e2dcd008-4d51-47dc-adaa-8d2b8d0dd785",
					EmailAddress: "xFowler@Demimbu.mil",
					PhoneNumber:  "125-00-48",
					Company:      "InnoZ",
					Brand:        "Livefish",
					Product:      "Output Viewer",
					Color:        "Turquoise",
					Hexcolor:     "#7d939e",
					Latitude:     84.324295,
					Longitude:    -99.350929,
					Floatn:       1.6909,
					Boolean:      true,
					SomeOneof: &Complete_OrCouldBe{
						OrCouldBe: &SubMessage{
							Id: "c8164fb3-2b06-494b-a40d-0ff64dcb8e86",
						},
					},
					AnotherOneof: &Complete_This{
						This: "est",
					},
				},
				&Complete{
					Id:          "bebf601b-0dc0-4ffb-a7e6-525c2f91f6b0",
					Status:      "complete",
					Description: "consequatur exercitationem consequatur velit minus nihil cupiditate qui doloremque et. tempore blanditiis nesciunt consequatur ut. necessitatibus aliquam nulla quibusdam voluptates rerum dolore id. excepturi ut placeat blanditiis repellendus accusamus. qui suscipit error molestiae eius quis incidunt.",
					Email:       "quaerat_tenetur@Edgewire.com",
					NotEmail:    "et",
					Phone:       "4-817-304-03-60",
					Word: []string{
						"ea",
						"et",
						"provident",
						"aut",
						"quis",
						"culpa",
						"vel",
						"voluptatibus",
					},
					Url:          "https://aimbu.gov/illum",
					SingleNumber: 329,
					RepeatedNumber: []int64{
						721,
						327,
						904,
					},
					Lat:       -19,
					Lng:       24,
					Words:     "enim corporis voluptatem",
					Wordsn:    "vitae quae sint soluta delectus atque aut harum aut nihil",
					Intn:      2,
					Fullname:  "Stephen Boyd",
					Firstname: "Pamela",
					Lastname:  "Austin",
					Paragraph: "facilis eum vel. vero omnis cupiditate et. voluptas quidem necessitatibus et rerum. amet odio non in aut tempore expedita labore error ullam.",
					Paragraphs: "quam veritatis hic earum ab nobis illo. dolore quasi tempora non ad. esse enim rem officiis nostrum optio.	quas neque et quasi quia recusandae at. et id pariatur neque iure omnis impedit. non ipsum quas sed aut aspernatur debitis nihil corrupti quo. autem nihil odit quo non dolorem. a quisquam fugiat quis aut in voluptas tenetur. ex voluptas eos iure esse. cupiditate expedita quia laudantium perspiciatis corrupti. esse sunt illo rem ut eum. magni sit nihil dolorem. qui sint non est.	asperiores vel id vero corrupti dignissimos. esse labore maiores ut fuga. consequuntur nobis aut doloribus porro ex consequatur. qui rerum aut corporis. natus quibusdam animi delectus eos alias autem distinctio. voluptas sint soluta quod ipsa et. sapiente et et totam et ab sed. dolorum sunt eum dolor alias distinctio sit saepe aut doloremque. aliquid omnis et temporibus libero eius pariatur.	cupiditate sunt modi possimus et deserunt. odit ut tempore. qui eligendi in eius. omnis architecto rerum quos asperiores unde asperiores rerum. ipsa et quidem consequatur quod. natus quibusdam sit qui ut et eaque maxime. nam vero minus asperiores provident incidunt quaerat. sed sit explicabo voluptatem fugiat explicabo voluptatem est. facilis illo deserunt accusantium eum!	totam distinctio repellendus ipsam facilis in sint eius.",
					Paragraphsn: "non nihil omnis cupiditate vel. corporis delectus molestiae voluptatem aliquid vero nulla. sed ducimus et incidunt. consectetur inventore sed ab nostrum asperiores sunt quis cum est. reprehenderit id magni porro necessitatibus quis voluptatem.	eum quos minus culpa rerum. rem vitae architecto quo velit aliquam quibusdam at reiciendis in. voluptatibus excepturi doloremque rerum magnam id aliquid dolorem. numquam perspiciatis excepturi quae! id cupiditate quia quasi. consectetur dolor alias similique placeat. eveniet quia corporis consequatur aspernatur quia omnis qui quia aut! cupiditate officia quod quam animi amet et.",
					Uuid:         "394ad4c2-4420-44be-b595-a83c11a17ecc",
					EmailAddress: "rSmith@Dynabox.info",
					PhoneNumber:  "4-815-173-88-16",
					Company:      "Yakijo",
					Brand:        "Meevee",
					Product:      "Auto HD Kit",
					Color:        "Mauv",
					Hexcolor:     "#3b3f44",
					Latitude:     50.203964,
					Longitude:    -89.096451,
					Floatn:       1.7203,
					Boolean:      true,
					AnotherOneof: &Complete_That{
						That: "ut sunt error praesentium amet debitis unde officia! in explicabo est sed deserunt autem. officiis id aut vel qui et nam maiores. qui sint sit nam pariatur ut non reprehenderit fuga. ut et sit deserunt quia. occaecati natus doloremque qui rem ea. sint reprehenderit veniam laboriosam. dolore voluptatem fugit quam eum. magni iste deserunt quaerat. consequatur voluptas et rerum nam pariatur.",
					},
					SomeOneof: &Complete_OrCouldBe{
						OrCouldBe: &SubMessage{
							Id: "0ea57d01-a0ca-493b-9d28-beb2e728591c",
						},
					},
				},
				&Complete{
					Id:          "f267332d-0d9a-4220-b055-63b661b600db",
					Status:      "in progress",
					Description: "voluptatum et deleniti. ipsam sit nihil. modi rerum est eos voluptatum! voluptas in qui dolorem et exercitationem. et odio beatae ut et nemo et omnis quae. inventore quae sed. optio voluptate amet quia nulla autem fuga assumenda veniam. placeat quo et consequatur culpa repellat quod consequatur. et sit doloremque exercitationem laborum qui quia sed.",
					Email:       "dolorem_maiores_corrupti@Skipstorm.info",
					NotEmail:    "qui",
					Phone:       "403-58-17",
					Word: []string{
						"minus",
						"et",
						"iste",
						"dolorem",
						"earum",
						"voluptatem",
						"voluptate",
						"illo",
					},
					Url:          "https://browseblab.edu/optio/consectetur/sit/ipsum",
					SingleNumber: 913,
					RepeatedNumber: []int64{
						123,
						111,
						61,
					},
					Lat:       14,
					Lng:       -92,
					Words:     "labore consectetur qui aut",
					Wordsn:    "sint et nemo voluptas repudiandae culpa perspiciatis voluptatem possimus nostrum",
					Intn:      2,
					Fullname:  "Terry Simpson",
					Firstname: "Antonio",
					Lastname:  "Hansen",
					Paragraph: "voluptas officiis facilis veniam omnis. culpa cum voluptatem officiis velit facilis. nihil quia voluptatem beatae ut laborum. ab explicabo animi consequatur iure culpa architecto iusto aut. laudantium voluptate voluptatem ut architecto omnis aut consequatur. cupiditate harum delectus quas deserunt autem suscipit. repellat aut voluptatem est autem animi veritatis quis. et facilis sit quia voluptas dolorem et dignissimos occaecati. sed et optio iure modi.",
					Paragraphs: "qui quos qui sit ratione eaque est!	ex ipsa quo sed. quo necessitatibus ratione quaerat voluptas. molestiae et et. architecto praesentium aut at exercitationem.	qui consequuntur recusandae. quaerat optio unde velit facilis labore libero! quia rerum ea nisi repudiandae consequuntur. dolores est aut itaque itaque consequuntur. aut facilis voluptatem ut qui dolores ea aut voluptates!",
					Paragraphsn: "et eius rem mollitia deserunt deserunt ipsam nihil dolorum. excepturi sit dignissimos ad vero. inventore et consequatur ipsa ipsam illum laudantium. est assumenda repellendus. voluptatem aperiam eaque et quis. eaque tempore velit non fugiat rerum magni.	quae totam illum iure est sunt. nostrum quaerat maiores exercitationem aut aut. aspernatur ducimus laborum magnam in beatae ut. assumenda autem doloremque voluptates ut quia veniam facilis. quos pariatur enim et molestiae qui. hic hic aut delectus culpa natus vitae. error eveniet quo saepe nobis. corporis eos sed est unde.",
					Uuid:         "8b081844-6397-4c14-a8a4-e6a1b79947ac",
					EmailAddress: "oScott@Skippad.com",
					PhoneNumber:  "2-173-013-53-87",
					Company:      "Wikibox",
					Brand:        "Flashspan",
					Product:      "Direct Input Mount",
					Color:        "Blue",
					Hexcolor:     "#10a2c2",
					Latitude:     87.379486,
					Longitude:    -155.610611,
					Floatn:       1.599,
					Boolean:      false,
					SomeOneof: &Complete_CouldBe{
						CouldBe: "eum",
					},
					AnotherOneof: &Complete_That{
						That: "accusantium cumque ut aut nobis quod. voluptates iste laboriosam error iusto perferendis sed voluptatem. veritatis molestias sint est excepturi aut.",
					},
				},
			},
		}
	return res, nil
}