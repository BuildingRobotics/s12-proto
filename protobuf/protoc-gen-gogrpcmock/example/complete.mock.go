// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: complete.proto

package s12_complete

import (
	context "context"
	fmt "fmt"
	_ "github.com/BuildingRobotics/s12-proto/protobuf/s12proto"
	proto "github.com/gogo/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type MyCompleteServiceMock struct{}

func (m *MyCompleteServiceMock) GetMessage(ctx context.Context, req *MsgRequest) (*MsgResponse, error) {
	res :=
		&MsgResponse{
			Id: "774f2362-24ec-4b5c-85e9-d24a54b3d453",
			Completed: []*Complete{
				&Complete{
					Id:          "698b7c58-2e8d-4f44-9d7c-f4389459217d",
					Status:      "complete",
					Description: "molestias distinctio blanditiis aut accusamus voluptas. eum autem qui iste beatae. qui fuga hic numquam animi ut ipsum sit. tempora sapiente ut fugiat voluptas exercitationem porro explicabo. aut tenetur recusandae voluptatum eos deserunt. consequatur eaque possimus. quia et porro laborum quis. rem deserunt expedita non! ad illo omnis libero labore et. sint rem dolore sint modi dolorum aut deleniti animi.",
					Email:       "cumque_totam_vitae@Topicstorm.mil",
					NotEmail:    "non",
					Phone:       "4-377-696-71-04",
					Word: []string{
						"sit",
						"quaerat",
						"iste",
					},
					Url:          "https://zoomzone.name/autem/alias",
					SingleNumber: 829,
					RepeatedNumber: []int64{
						517,
						985,
						220,
					},
					Lat:        23,
					Lng:        72,
					Words:      "labore alias voluptatem soluta",
					Wordsn:     "consequuntur fuga distinctio assumenda cupiditate nihil distinctio nam deserunt optio",
					Intn:       0,
					Fullname:   "Edward Hernandez",
					Firstname:  "Melissa",
					Lastname:   "Scott",
					Paragraph:  "nihil illo quia qui. voluptatibus fugit incidunt cum aliquam eius fugit.",
					Paragraphs: "atque sapiente alias aliquid quae earum architecto. ut cum quis qui id et. ut ut ducimus eaque sed ducimus ipsam aliquid! earum nobis ex voluptas! fugit voluptas consequatur et quis voluptas! earum quidem saepe quod consequatur eveniet.",
					Paragraphsn: "eius nisi quidem natus atque non possimus. mollitia amet et quis dignissimos dolores.	est cupiditate non molestias esse ut.",
					Uuid:         "5d69cb35-d0d8-4006-a181-79238fe980fb",
					EmailAddress: "dolor_et@Meejo.gov",
					PhoneNumber:  "238-82-09",
					Company:      "Centidel",
					Brand:        "Skalith",
					Product:      "Portable System",
					Color:        "Mauv",
					Hexcolor:     "#0000d0",
					Latitude:     33.532104,
					Longitude:    -77.150032,
					Floatn:       1.8799,
					Boolean:      true,
					SomeOneof: &Complete_OrCouldBe{
						OrCouldBe: &SubMessage{
							Id: "dc9fe12e-87ff-4553-9466-fad3f19ded56",
						},
					},
					AnotherOneof: &Complete_That{
						That: "et harum ut amet aperiam. facere ratione cupiditate excepturi et dolore quo. minus alias quas ea ipsam laborum ea vero aut quia esse. quam non sunt dignissimos assumenda ex cupiditate ut! error eos ut omnis!",
					},
				},
				&Complete{
					Id:          "22a4d46b-74c1-4d82-9ffe-3f1c43db10bc",
					Status:      "in progress",
					Description: "maxime nostrum unde dolorem sit. non qui debitis et. veritatis earum rerum. officiis nostrum incidunt nihil eum commodi et! tenetur dolores excepturi eos iure nulla. quaerat voluptatibus omnis qui natus aut dolorum soluta. ratione nam ut est vel sint. rerum cupiditate voluptatibus libero ipsa. enim dicta facilis et fugiat.",
					Email:       "sRiley@Zoozzy.org",
					NotEmail:    "rerum",
					Phone:       "629-84-20",
					Word: []string{
						"quia",
						"laborum",
						"dolor",
					},
					Url:          "https://devbug.org/earum/et/et",
					SingleNumber: 605,
					RepeatedNumber: []int64{
						623,
						838,
						527,
					},
					Lat:       -4,
					Lng:       -142,
					Words:     "maiores saepe assumenda eius corrupti",
					Wordsn:    "ut aut quis delectus maxime possimus in eveniet et sit",
					Intn:      4,
					Fullname:  "Alice Wilson",
					Firstname: "Ralph",
					Lastname:  "Cox",
					Paragraph: "sint corrupti itaque laboriosam quia. debitis hic velit exercitationem voluptatem et molestiae. vitae recusandae repellendus qui nesciunt similique sunt. impedit nostrum eos illum et quam. distinctio impedit ut doloribus. est nulla voluptas velit nihil error sit. libero et et consequatur. cum et rerum est nesciunt inventore voluptas. culpa fuga voluptatem.",
					Paragraphs: "accusantium saepe exercitationem. sed sed ea suscipit ipsa et. et et neque rerum. nam error excepturi sed sint adipisci incidunt est. sed eaque ut et sunt officia illo qui enim!	non voluptas distinctio. iusto sint molestias minus at.",
					Paragraphsn: "alias ex distinctio qui! qui voluptatem sit est omnis aliquid. voluptates velit a cupiditate autem quia. explicabo corporis veritatis mollitia. culpa ipsa est vero id est dolorem. veritatis libero reprehenderit voluptatibus voluptatibus quas laudantium neque. quia impedit tenetur ea nam est! voluptatum distinctio magni. consequatur perferendis doloribus et labore. assumenda et omnis quo est repellendus pariatur.	architecto excepturi expedita perferendis! tenetur quisquam nemo sapiente vel consequatur eaque reiciendis. adipisci et vel quis veritatis. est id sunt non adipisci. unde ut et est. nam similique impedit voluptatem et nesciunt quam!",
					Uuid:         "229a358c-884d-4009-8576-5e6bad00913e",
					EmailAddress: "BrandonPierce@Fivechat.gov",
					PhoneNumber:  "4-468-963-10-05",
					Company:      "Vimbo",
					Brand:        "Flashdog",
					Product:      "GPS System",
					Color:        "Green",
					Hexcolor:     "#a81466",
					Latitude:     -64.145096,
					Longitude:    45.999023,
					Floatn:       0.3305,
					Boolean:      true,
					SomeOneof: &Complete_OrCouldBe{
						OrCouldBe: &SubMessage{
							Id: "1b6dd0db-a709-43ca-8523-0d081b2328cd",
						},
					},
					AnotherOneof: &Complete_That{
						That: "voluptates at amet atque tempore. nesciunt optio error commodi cupiditate. aspernatur quam quis sunt qui libero excepturi similique! sit et doloremque sunt! earum quam aperiam ut perferendis eaque qui! enim veniam occaecati.",
					},
				},
				&Complete{
					Id:          "b756128d-a5e3-4c34-ae3b-b4748bf05431",
					Status:      "in progress",
					Description: "fugiat voluptatem excepturi. aperiam quas totam asperiores fuga.",
					Email:       "magnam_fuga_fuga@Riffwire.gov",
					NotEmail:    "itaque",
					Phone:       "7-386-011-35-66",
					Word: []string{
						"sit",
						"ullam",
						"officiis",
					},
					Url:          "https://skiptube.org/ut/culpa/non/voluptate",
					SingleNumber: 848,
					RepeatedNumber: []int64{
						680,
						486,
						138,
					},
					Lat:       -5,
					Lng:       -56,
					Words:     "qui earum rerum",
					Wordsn:    "omnis et velit qui vel odio perferendis non qui iusto",
					Intn:      0,
					Fullname:  "James Hicks",
					Firstname: "Christopher",
					Lastname:  "Arnold",
					Paragraph: "maiores sed et quo quos. voluptatem eius tenetur possimus nam aut. sint sit minima temporibus cumque modi quis dolorem.",
					Paragraphs: "aut dolorum nesciunt voluptatem enim autem eius. laborum architecto magni ea ab et. molestiae vitae incidunt soluta ut quas incidunt expedita.	voluptatem qui velit unde. qui quam corrupti placeat! in autem ut eius quibusdam hic. repellat aliquam dolores. architecto qui atque aut praesentium nesciunt quod.",
					Paragraphsn: "labore cumque quisquam est accusamus rem. enim laudantium consectetur est quia! harum sunt et. vero culpa quidem eveniet est suscipit nisi nobis.	dicta aut ratione occaecati eos facere officia.",
					Uuid:         "d0648f73-abed-4df5-bcc5-52ec3e7793ea",
					EmailAddress: "mollitia@Devpoint.org",
					PhoneNumber:  "6-592-021-89-80",
					Company:      "Zazio",
					Brand:        "Talane",
					Product:      "Performance Output Kit",
					Color:        "Yellow",
					Hexcolor:     "#8c2c8a",
					Latitude:     -57.566391,
					Longitude:    -97.010399,
					Floatn:       2.7307,
					Boolean:      true,
					AnotherOneof: &Complete_This{
						This: "eligendi",
					},
					SomeOneof: &Complete_CouldBe{
						CouldBe: "non",
					},
				},
			},
		}
	return res, nil
}
