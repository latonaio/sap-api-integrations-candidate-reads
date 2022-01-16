# sap-api-integrations-candidate-reads
sap-api-integrations-candidate-reads は、外部システム(特にエッジコンピューティング環境)をSAPと統合することを目的に、SAP API で 候補者 データを取得するマイクロサービスです。    
sap-api-integrations-candidate-reads には、サンプルのAPI Json フォーマットが含まれています。   
sap-api-integrations-candidate-reads は、SAP SuccessFactors API の利用を前提としています。  
https://api.sap.com/api/RCMCandidate/overview   

## 動作環境  
sap-api-integrations-candidate-reads は、主にエッジコンピューティング環境における動作にフォーカスしています。  
使用する際は、事前に下記の通り エッジコンピューティングの動作環境（推奨/必須）を用意してください。  
・ エッジ Kubernetes （推奨）    
・ AION のリソース （推奨)    
・ OS: LinuxOS （必須）    
・ CPU: ARM/AMD/Intel（いずれか必須）    

## クラウド環境での利用
sap-api-integrations-candidate-reads は、外部システムがクラウド環境である場合にSAPと統合するときにおいても、利用可能なように設計されています。  

## 本レポジトリ が 対応する API サービス
sap-api-integrations-candidate-reads が対応する APIサービス は、次のものです。

* APIサービス概要説明 URL: https://api.sap.com/api/RCMCandidate/overview    
* APIサービス名(=baseURL): https://sandbox.api.sap.com/successfactors/

## 本レポジトリ に 含まれる API名
sap-api-integrations-candidate-reads には、次の API をコールするためのリソースが含まれています。  

* Candidate（候補者 - ヘッダ）※候補者の詳細データを取得するために、Languages、Education、Certificates、OutsideWorkExperience、JobsApplied、Resumeと合わせて取得されます。
* CandidateBackground_Languages（候補者 - 言語能力）
* CandidateBackground_Education（候補者 - 学歴）
* CandidateBackground_Certificates（候補者 - 資格）
* CandidateBackground_OutsideWorkExperience（候補者 - 職歴）
* Languages（候補者 - 言語能力 ※To）
* Education（候補者 - 学歴 ※To）
* Certificates（候補者 - 資格 ※To）
* OutsideWorkExperience（候補者 - 職歴 ※To）
* JobsApplied（候補者 - 応募求人 ※To）
* Resume（候補者 - レジュメ ※To）
* State（候補者 - 状況 ※To）

## API への 値入力条件 の 初期値
sap-api-integrations-candidate-readsにおいて、API への値入力条件の初期値は、入力ファイルレイアウトの種別毎に、次の通りとなっています。  

### SDC レイアウト

* inoutSDC.Candidate.CandidateID（候補者ID）
* inoutSDC.Candidate.FirstName（ファーストネーム）
* inoutSDC.Candidate.LastName（ラストネーム）

## SAP API Bussiness Hub の API の選択的コール

Latona および AION の SAP 関連リソースでは、Inputs フォルダ下の sample.json の accepter に取得したいデータの種別（＝APIの種別）を入力し、指定することができます。  
なお、同 accepter にAll(もしくは空白)の値を入力することで、全データ（＝全APIの種別）をまとめて取得することができます。  

* sample.jsonの記載例(1)  

accepter において 下記の例のように、データの種別（＝APIの種別）を指定します。  
ここでは、"Role" が指定されています。    
  
```
	"api_schema": "Candidate",
	"accepter": ["Header"],
	"candidate_id": "1283",
	"deleted": false
```
  
* 全データを取得する際のsample.jsonの記載例(2)  

全データを取得する場合、sample.json は以下のように記載します。  

```
	"api_schema": "Candidate",
	"accepter": ["All"],
	"candidate_id": "1283",
	"deleted": false
```

## 指定されたデータ種別のコール

accepter における データ種別 の指定に基づいて SAP_API_Caller 内の caller.go で API がコールされます。  
caller.go の func() 毎 の 以下の箇所が、指定された API をコールするソースコードです。  

```
func (c *SAPAPICaller) AsyncGetCandidate(candidateID, firstName, lastName string, accepter []string) {
	wg := &sync.WaitGroup{}
	wg.Add(len(accepter))
	for _, fn := range accepter {
		switch fn {
		case "Header":
			func() {
				c.Header(candidateID)
				wg.Done()
			}()
		case "Languages":
			func() {
				c.Languages(candidateID)
				wg.Done()
			}()
		case "Education":
			func() {
				c.Education(candidateID)
				wg.Done()
			}()
		case "Certificates":
			func() {
				c.Certificates(candidateID)
				wg.Done()
			}()
		case "OutsideWorkExperience":
			func() {
				c.OutsideWorkExperience(candidateID)
				wg.Done()
			}()
		case "CandidateByName":
			func() {
				c.CandidateByName(firstName, lastName)
				wg.Done()
			}()
		default:
			wg.Done()
		}
	}

	wg.Wait()
}
```
## Output  
本マイクロサービスでは、[golang-logging-library](https://github.com/latonaio/golang-logging-library) により、以下のようなデータがJSON形式で出力されます。  
以下の sample.json の例は、候補者 の ヘッダデータ が取得された結果の JSON の例です。  
以下の項目のうち、"CandidateID" ～ "CellPhone" は、/SAP_API_Output_Formatter/type.go 内 の Type Header {} による出力結果です。"cursor" ～ "time"は、golang-logging-library による 定型フォーマットの出力結果です。  

```
{
	"cursor": "/Users/latona2/bitbucket/sap-api-integrations-candidate-reads/SAP_API_Caller/caller.go#L63",
	"function": "sap-api-integrations-candidate-reads/SAP_API_Caller.(*SAPAPICaller).Header",
	"level": "INFO",
	"message": [
		{
			"CandidateID": "1283",
			"Country": "Japan",
			"PartnerMemberID": "",
			"LastLoginDateTime": "/Date(1446091083000+0000)/",
			"LastModifiedDateTime": "/Date(1446091101000+0000)/",
			"AnonymizedDateTime": "",
			"Anonymized": "Non-anonymized",
			"PublicIntranet": false,
			"UsersSysID": "",
			"ExternalCandidate": true,
			"PrimaryEmail": "mfukui@email.com",
			"CreationDateTime": "/Date(1446009621000+0000)/",
			"Zip": "5420047",
			"HomePhone": "",
			"FirstName": "正治",
			"PrivacyAcceptDateTime": "",
			"CurrentTitle": "",
			"ConsentToMarketing": "",
			"AgreeToPrivacyStatement": "false",
			"DateofAvailability": "",
			"LastName": "福井",
			"City": "大阪",
			"DataPrivacyID": "",
			"VisibilityOption": false,
			"Address": "2-4-6",
			"CandidateLocale": "en_GB",
			"Address2": "中央区",
			"ContactEmail": "test@abc.com",
			"ShareProfile": "Any company recruiter worldwide",
			"PartnerSource": "0",
			"MiddleName": "",
			"CellPhone": "+81+81673748273"
		}
	],
	"time": "2022-01-10T13:21:19.551+09:00"
}
```
