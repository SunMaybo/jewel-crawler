package parser

import (
	"fmt"
	"testing"
)

var content = `

<!DOCTYPE html>
<html xmlns="https://www.w3.org/1999/xhtml" lang="en">
<!-- ChristiesPlus -->
<head>
	

  <!-- OneTrust Cookies Consent Notice start -->
  
  <script src="https://cdn.cookielaw.org/consent/1d0e237c-19a3-4077-bea6-ecb56a0c3401.js" type="text/javascript" data-document-language="true" charset="UTF-8"></script>
  
  <script type="text/javascript">
    function OptanonWrapper() {
      Optanon.InsertScript('/js/performance-cookies.js', 'head', null, null, 2);
      Optanon.InsertScript('/js/targeting-cookies.js', 'head', null, null, 4);
    }
  </script>
  <!-- OneTrust Cookies Consent Notice end -->
  

  <style media="screen">
    .christies-header{display:block;height:54px}@media (min-width:768px){.christies-header{height:58px}}@media (min-width:1280px){.christies-header{height:109px}}.lazyload{opacity:.5;transition:opacity .15s ease-in}.lazyload:not(.loaded){height:100%;width:100%}.lazyload.loaded{opacity:1}.loader{position:fixed;top:0;left:0;width:100vw;height:100vh;background:rgba(0,0,0,.15);z-index:9999999}.loader--inner{height:60px;width:60px;top:0;right:0;bottom:0;left:0;position:fixed;margin:auto;border:6px solid #fff;border-right-color:#b30900;border-top-color:#b30900;border-radius:100%;animation:spin .8s infinite linear;z-index:1200}.comp-loader{position:absolute;top:0;left:0;width:100%;min-height:250px;background:rgba(0,0,0,0);z-index:99}.comp-loader--inner{height:60px;width:60px;top:0;right:0;bottom:0;left:0;position:absolute;margin:auto;border:6px solid #fff;border-right-color:#b30900;border-top-color:#b30900;border-radius:100%;animation:spin .8s infinite linear;z-index:12}@keyframes spin{from{transform:rotate(0)}to{transform:rotate(359deg)}}
  </style>
  <link rel="stylesheet" href="/app_build/ccHeaderFooter.css" />

  <script type="text/javascript">
    //define window.tempBind - this is used, in conjunction with the angular-custom.js file, to prevent scriptaculous and legacy prototype js from breaking angular
    window.tempBind = Function.prototype.bind;

    function getParameterByName(StrURL, name) {
      name = name.replace(/[\[]/, "\\[").replace(/[\]]/, "\\]");
      var regex = new RegExp("[\\?&]" + name + "=([^&#]*)"),
        results = regex.exec(StrURL.toLocaleLowerCase());
      return results == null ? "" : decodeURIComponent(results[1].replace(/\+/g, " "));
    }

    function getLanguage() {
      var pageLang = "";
      var loc = window.location.href.toLowerCase();
      var patt = /departments\/.*-(\d{1,3})-(\d{1,2})\.aspx/;
      var lang;
      var IsDepartment = patt.test(loc);
      //If the language is always coming from the backend, what's the purpose of the following if condition? What is a use case for this?
      if (pageLang == null || pageLang == "") {
        if (loc.indexOf("/zh-cn") > -1 || getParameterByName(loc, "sc_lang") == "zh-cn" || getParameterByName(loc, "lid") == "4" || getParameterByName(loc, "languageid") == "4" || (IsDepartment && loc.indexOf("-4.aspx") > -1)) {
          lang = "zh-cn";
        } else if (loc.indexOf("/zh") > -1 || getParameterByName(loc, "sc_lang") == "zh" || getParameterByName(loc, "lid") == "3" || getParameterByName(loc, "languageid") == "3" || (IsDepartment && loc.indexOf("-3.aspx") > -1)) {
          lang = "zh";
        } else {
          lang = "en";
        }
      } else {
        lang = pageLang;
      }
      return lang;

    }

    var _val = getLanguage();
    /*var d = new Date();
    d.setTime(d.getTime() + (1 * 24 * 60 * 60 * 1000));
    var expires = "expires=" + d.toUTCString();*/
    document.cookie = "CurrentLanguage=" + _val + ";Domain=.christies.com;path=/";

    function setCookie(name, value, expires, path, domain, secure) {
      cookieStr = name + "=" + escape(value) + "; ";

      if (expires) {
        expires = setExpiration(expires);
        cookieStr += "expires=" + expires + "; ";
      }
      if (path) {
        cookieStr += "path=" + path + "; ";
      }
      if (domain) {
        cookieStr += "domain=" + domain + "; ";
      }
      if (secure) {
        cookieStr += "secure; ";
      }

      document.cookie = cookieStr;
    }

    function setExpiration(cookieLife) {
      var today = new Date();
      var expr = new Date(today.getTime() + cookieLife * 24 * 60 * 60 * 1000);
      return expr.toGMTString();
    }

    setCookie("CurrentLanguage", "", "-100", "/", "www.christies.com", false);
  </script>

  <meta charset="UTF-8" />
  <meta http-equiv="Content-Type" content="text/html; charset=UTF-8" />
  <meta http-equiv="X-UA-Compatible" content="IE=edge" />
  <meta http-equiv="Cache-Control" content="no-cache" />
  <meta name="viewport" content="width=device-width, initial-scale=1.0, maximum-scale=1.0, user-scalable=0" />
  <meta name="apple-mobile-web-app-capable" content="yes" />
  <meta name="google-site-verification" content="2pm6xoJkRM_9TezjADtdaKy43hAcMXqYJWwwxc1viVY" />
  

    <link rel="canonical" href="https://www.christies.com/Jewels-Online-28905.aspx" />
  	
  <!-- SLP	-->
  <link href="/Build/css/v3/slp-bundle.min.css?180720201156" rel="stylesheet" />
  

  <link href="/static/images/favicon.ico?v=1" rel="shortcut icon" type="image/x-icon" />

  <!-- Can this be moved to the footer without any impact? -->
  <script type="text/javascript">
    var sOmnitureEnvironment = 'christiesprod';
    var APIKey_caption = "APIKey";
    var mychristies_key = '';
    var mychristies_url = 'https://www.christies.com/ChristiesAPIServices/mychristiesapi/api/MyChristies/';
    var displayFsu = true;
    var GEOCountryCode ='VE'
  </script>


  
  <meta name="og:title" property="og:title" content="Jewels Online" />
  <meta name="og:description" property="og:description" content="Bid in-person or online for the upcoming auction:Jewels Online on 22 July - 7 August 2020 Online at Christies.com" />
  <meta name="og:image" property="og:image" content="https://www.christies.com/images/generic_logo_no_catcover.png" />
  <meta name="og:url" property="og:url" content="" />
  <meta name="og:type" property="og:type" content="website" />
  <meta name="description" content="Bid in-person or online for the upcoming auction:Jewels Online on 22 July - 7 August 2020 Online at Christies.com">
 

  <meta name="google-site-verification" content="2pm6xoJkRM_9TezjADtdaKy43hAcMXqYJWwwxc1viVY" />

  
    <meta name="apple-itunes-app" content="app-id=322000464,app-argument=christies://christies.app/sale?saleID=28905" />
  
  
  <link rel="canonical" href="https://www.christies.com/Jewels-Online-28905.aspx" />
  

  <meta name="google-site-verification" content="2pm6xoJkRM_9TezjADtdaKy43hAcMXqYJWwwxc1viVY" />

  
  <link rel="canonical" href="https://www.christies.com/Jewels-Online-28905.aspx" />
  <meta id="meta_keywords" name="keywords" content="Christies,auction" />
  <meta id="meta_description" name="description" />
 
  


  
  <title>
    Jewels Online | Christie's
      
  </title>
  
  <script src="/Build/scripts/v3/global-bundle.min.js?180720201156"></script>
</head>
<body id="index" class='index  English'>
  <!-- Page body -->

  <header>
    <div class='christies-header'>
      <cc-global-header></cc-global-header>
    </div>
  </header>
  
  
<!-- ===========================================
    This should be the first thing just after the  Body Tag
        OMNITURE CODE SECTION-A BEGIN
     ===========================================-->

<script type="text/javascript" language="javascript" src="//assets.adobedtm.com/4f105c1434ad/6ef91fc289ec/launch-8bada67044f4.min.js"></script>

<script type="text/javascript" language="javascript">var s = s || {};</script>

<script  type='text/javascript' language='JavaScript'><!--
s.prop37 = window.location.href;
var _userLanguage = "English";
   var _userLanguageID = "1";
 var _OmnitureEnvironment = "christiesprod";
//--></script>
<!-- ===========================================
    OMNITURE CODE SECTION-A END
     ===========================================-->

  <form method="post" action="/salelanding/index.aspx?salenumber=19601&amp;saleroomcode=NYR&amp;lid=1" id="mainform">
<div class="aspNetHidden">
<input type="hidden" name="__EVENTTARGET" id="__EVENTTARGET" value="" />
<input type="hidden" name="__EVENTARGUMENT" id="__EVENTARGUMENT" value="" />
<input type="hidden" name="__VIEWSTATE" id="__VIEWSTATE" value="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" />
</div>

<script type="text/javascript">
//<![CDATA[
var theForm = document.forms['mainform'];
if (!theForm) {
    theForm = document.mainform;
}
function __doPostBack(eventTarget, eventArgument) {
    if (!theForm.onsubmit || (theForm.onsubmit() != false)) {
        theForm.__EVENTTARGET.value = eventTarget;
        theForm.__EVENTARGUMENT.value = eventArgument;
        theForm.submit();
    }
}
//]]>
</script>


<div class="aspNetHidden">

	<input type="hidden" name="__VIEWSTATEGENERATOR" id="__VIEWSTATEGENERATOR" value="8C6CECF3" />
	<input type="hidden" name="__EVENTVALIDATION" id="__EVENTVALIDATION" value="IAnxnD3Mh8DRim0iYkFLNVIEIe4gOccDGT+8iZDFahDDJxUdiOzUMEKHNwTya5ms3z22C5E3LtpjR9jIIvqkpbzHQ+JcV1DuVBasgfbgd/g/4Gq9bz5h5JmbT/HNGrhLmtGEkfJroP1f7H9i5GsAsuHq0tycEt5SBACvgjWvI1mkcbrGkoyTVeF/zYpNHRKwiEa4yUpZBihaR8vEIIoV89nXIP0=" />
</div>
    

<script type="text/javascript">
    var LotFinderAPI = "https://www.christies.com/interfaces/LotFinderAPI/SearchResults.asmx/";
    var FollowLotAPI = "https://www.christies.com/webservices/mobile/Followlot.asmx/";
    var FollowSaleAPI = "https://www.christies.com/webservices/mobile/FollowSale.asmx/";
    var SecureRootPath = "https://www.christies.com/";
    var YouMayAlsoLikeMaxLimit = "150";
    var RecentKeywordSearchLimit = "5";
    var strRecentSearch = "RECENT SEARCH";
    var strPastLots = "Sold Lots";
    var strUpcomingLots = "Upcoming Lots";
    var strMobilSearchPlaceHolder = "Search Keyword";
    var IsSaleover = "False";
    var hdnClientGUID = "main_center_0_hdnClientGUID";
    var hdnLangID = "main_center_0_hdnLangID";
    var hdnGeoCountryCode = "main_center_0_hdnGeoCountryCode";
    var sMaptype = "google";
	var SaleLandingAPIKey = "";
    var SaleLandingAPI = "https://www.christies.com/ChristiesAPIServices/SaleLandingAPI/api/SaleLanding/";
    var AkamaiCaching = "false";
	var intSaleID = '28905';
    var sSaleNumber='19601';
	var sSaleRoomCode='NYR';
	var bIsSaleover='False';
	var sSaleTYPE='3';
    var bIsAsiaMobilePage = 'False';
    var preview = 'n'
	var _RelatedFeatureUrl = "/AjaxPages/SaleLanding/RelatedFeatures.aspx?intsaleid=" + intSaleID + "&lid=" + 1 + "&SaleNumber=" + sSaleNumber + "&SaleRoomCode=" + sSaleRoomCode + "&preview=" + preview;
	var _SaleInfoUrl = "/AjaxPages/SaleLanding/SaleInformation.aspx?intsaleid=" + intSaleID + "&lid=" + 1 + "&SaleNumber=" + sSaleNumber + "&SaleRoomCode=" + sSaleRoomCode + "&preview=" + preview;
	var _ContactDept = "/AjaxPages/SaleLanding/ContactDepartment.aspx?intsaleid=" + intSaleID + "&lid=" + 1 + "&SaleNumber=" + sSaleNumber + "&SaleRoomCode=" + sSaleRoomCode + "&preview=" + preview;
	var _RelatedSalesEvents = "/AjaxPages/SaleLanding/RelatedSalesAndEvents.aspx?intsaleid=" + intSaleID + "&lid=" + 1 + "&SaleNumber=" + sSaleNumber + "&SaleRoomCode=" + sSaleRoomCode + "&preview=" + preview;
	var _UpcomingAuctions = "/AjaxPages/SaleLanding/UpcomingAuctions.aspx?intsaleid=" + intSaleID + "&lid=" + 1 + "&SaleNumber=" + sSaleNumber + "&SaleRoomCode=" + sSaleRoomCode + "&preview" + preview;
	var _BrowseLots = "/AjaxPages/SaleLanding/DisplayLotList.aspx?" + "salenumber=19601&saleroomcode=NYR&lid=1";
    var _featureContentUrl = "/AjaxPages/SaleLanding/FeatureContent.aspx?intsaleid=" + intSaleID + "&lid=" + 1 + "&SaleNumber=" + sSaleNumber + "&SaleRoomCode=" + sSaleRoomCode + "&preview=" + preview;

</script>

<input type="hidden" name="main_center_0$hdnLangID" id="main_center_0_hdnLangID" value="1" />
<input type="hidden" name="main_center_0$hdnGeoCountryCode" id="main_center_0_hdnGeoCountryCode" value="VE" />
<input type="hidden" name="main_center_0$hdnClientGUID" id="main_center_0_hdnClientGUID" />
<input type="hidden" name="main_center_0$hdnClientID" id="main_center_0_hdnClientID" value="0" />

<div class="salelanding">
    <ul class="container container-modules">
        <li>
        <div id="sale-header">
            
<script>
  var sOmnitureDepartment = 'Jewellery';
  var blnWineMsg = 'False';
  var btnExcel = "main_center_0_ctl00_btnExcel";
  var EditChristiesLive = 'Edit registration';
  var SignInToBid = 'Sign in to bid';
  var RegistertoBid = 'Register to bid';
  var ShopNow = 'Shop Now';
  var BidNow = 'Bid Now';
  var followLot = 'Follow';
  var unfollowLot = 'Unfollow';
  var bIsSaleAvailableforLotfinder = 'False';
  var bIsSaleLandingPage = true;
    
    

</script>

<script type="text/javascript">
    var ddlHKAMobileSales = 'main_center_0_ctl00_ddlHKAMobileSales';
 //--HB-345 
        var dtm = {
clientGuid: '',
cultureCode: 'en-us',
languageName: 'ENGLISH'
};
dtm['saleId'] = '28905';
dtm['saleTitle'] = 'Jewels Online';
dtm['saleNumber'] = '19601';
dtm['pageType'] = 'SLP';
dtm['saleType'] ='Jewellery';

        //--HB-345 
</script>

<div id="main_center_0_ctl00_dvmainheader">
  <div class="col-xs-12 no-padding slp slp-header" id="MainLotHeader">
    <div class="sale-header">
      <div class="col-xs-12 col-sm-3 col-md-3 col-lg-3 no-padding main-image-wrapper" id="MainSaleImage">
        <div class="image--container_landscape">
            <picture>
              <source media="(max-width: 320px)" data-srcset="https://www.christies.com/img/saleimages/NYR-19601-07222020-1.JPG?maxwidth=280&quality=70, https://www.christies.com/img/saleimages/NYR-19601-07222020-1.JPG?maxwidth=560&quality=70 2x">

              <source media="(max-width: 375px)" data-srcset="https://www.christies.com/img/saleimages/NYR-19601-07222020-1.JPG?maxwidth=335&quality=70, https://www.christies.com/img/saleimages/NYR-19601-07222020-1.JPG?maxwidth=670&quality=70 2x, https://www.christies.com/img/saleimages/NYR-19601-07222020-1.JPG?maxwidth=1005&quality=70 3x">

              <source media="(max-width: 414px)" data-srcset="https://www.christies.com/img/saleimages/NYR-19601-07222020-1.JPG?maxwidth=374&quality=70, https://www.christies.com/img/saleimages/NYR-19601-07222020-1.JPG?maxwidth=748&quality=70 2x, https://www.christies.com/img/saleimages/NYR-19601-07222020-1.JPG?maxwidth=1122&quality=70 3x">

              <source media="(max-width: 736px)" data-srcset="https://www.christies.com/img/saleimages/NYR-19601-07222020-1.JPG?maxwidth=696&quality=70, https://www.christies.com/img/saleimages/NYR-19601-07222020-1.JPG?maxwidth=1392&quality=70 2x, https://www.christies.com/img/saleimages/NYR-19601-07222020-1.JPG?maxwidth=2090&quality=70 3x">

              <source media="(max-width: 768px)" data-srcset="https://www.christies.com/img/saleimages/NYR-19601-07222020-1.JPG?maxwidth=182&quality=70, https://www.christies.com/img/saleimages/NYR-19601-07222020-1.JPG?maxwidth=364&quality=70 2x, https://www.christies.com/img/saleimages/NYR-19601-07222020-1.JPG?maxwidth=546&quality=70 3x">

              <source media="(max-width: 1024px)" data-srcset="https://www.christies.com/img/saleimages/NYR-19601-07222020-1.JPG?maxwidth=246&quality=70, https://www.christies.com/img/saleimages/NYR-19601-07222020-1.JPG?maxwidth=492&quality=70 2x,https://www.christies.com/img/saleimages/NYR-19601-07222020-1.JPG?maxwidth=738&quality=70 3x">

              <source media="(min-width: 1025px)" data-srcset="https://www.christies.com/img/saleimages/NYR-19601-07222020-1.JPG?maxwidth=295&quality=70, https://www.christies.com/img/saleimages/NYR-19601-07222020-1.JPG?maxwidth=590&quality=70 2x, https://www.christies.com/img/saleimages/NYR-19601-07222020-1.JPG?maxwidth=885&quality=70 3x">
              <img class="image lazyload" data-src="https://www.christies.com/img/saleimages/NYR-19601-07222020-1.JPG?maxwidth=885&quality=70" alt="Jewels Online auction at Christies">
            </picture>
        </div>
      </div>

      <div class="col-xs-12 col-sm-6 col-md-7 col-lg-7 main-sale-details" id="MainSaleDetails">
        <span id="main_center_0_ctl00_lblSaleNumber" class="p--primary_small sale-reference">Sale 19601</span>
        <h2 class="heading-2_secondary sale-title"><span id="main_center_0_ctl00_lblSaleTitle">Jewels Online</span></h2>

        <div class="location-date--wrapper p--primary_large">
          <div>
            <span>
              Online</span>
            
            <span class="divider--left location-date">
              22 July - 7 August </span>
              
            <!-- <span class="hidden-xs hidden-sm hidden-md divider--left">Sale 19601</span> -->
          </div>
          <div class="hidden-lg location-date" style="display: none !Important">
            <!-- 22 July - 7 August , Online -->
            <span id="main_center_0_ctl00_spnSaleTimeCounter"></span>
            <span></span>
          </div>
        </div>

        

      
        <div class="no-padding upcoming-sale-materials" id="MainSaleMaterials">
          
          

          
          

          
          
          
          
          
             

           <div id="main_center_0_ctl00_dvTermsConditions" class="print-gallery pull-left">
            <div class="print termsConditions" id="dvTermsConditions">
              
            </div>
          </div>
        </div>

        
         
      </div>

      <div class="col-xs-12 col-sm-3 col-md-2 col-lg-2 no-padding hidden-print sale-btn-wrapper" id="MainSaleCTA">
        
        
        <div id="main_center_0_ctl00_dvBidToolTip" class="col-xs-12 bid-btn-wrapper no-padding heading-4">

          Bidding begins on 22 July

        </div>
        
        
      </div>
    </div>


    <div class="darkenBackground" id="divOverly" style="display: none"></div>


    <div id="divWineMessage" style="display: none" class="loginboxhover">
      

      <div class="innerPopupMessage">
        <span id="main_center_0_ctl00_lblWineMsg"></span>
        <br />
        <br />
        <div class="cta-wrapper">
          <a onclick="javascript:clsAlert();" id="main_center_0_ctl00_lnkAgree" class="col-xs-3 btn btn--primary btndownload pull-right" UseSubmitBehavior="False" href="javascript:__doPostBack(&#39;main_center_0$ctl00$lnkAgree&#39;,&#39;&#39;)"><span>Agree and Download</span>
          </a>

          <input type="submit" name="main_center_0$ctl00$btnDisAgree" value="Decline and Exit" id="main_center_0_ctl00_btnDisAgree" class="btn col-xs-3 btndownload pull-right" data-dismiss="modal" />
        </div>
      </div>
      <button type="button" class="close christies-icon_close" data-dismiss="modal" aria-hidden="false"></button>
    </div>

  </div>
</div>


        </div>
        
                <div id="sale-highlights">
            
<section class="panzoom--wrapper slp-zoom">
  <div class="loader">
    <div class="loader--inner"></div>
  </div>
  <div class="panzoom--container">
    <div class="panzoom--inner">
      <img class="panzoom" src="" />
    </div>
    <div class="buttons--container">
      <button class="zoom-out  slp-panzoom--icon ">
          <span class="icon-theme--light pull-left slp-panzoom--minus-icon">
                <svg class="icon svg--minus svg--icons slp-panzoom--minus-icon" xmlns="https://www.w3.org/2000/svg" xmlns:xlink="https://www.w3.org/1999/xlink" version="1.1" id="Layer_1" x="0px" y="0px" viewBox="0 0 32 32" xml:space="preserve">
					<path xmlns="http://www.w3.org/2000/svg" d="M29.7,4.6H2.3C1,4.6,0,3.5,0,2.3S1,0,2.3,0h27.4C31,0,32,1,32,2.3S31,4.6,29.7,4.6z"></path>
				</svg>
            </span>
      </button>
      <input type="range" class="zoom-range">
      <button class="zoom-in slp-panzoom--icon">
            <span class="icon-theme--light pull-right ">
                <svg class="icon svg--plus svg--icons" xmlns="https://www.w3.org/2000/svg" xmlns:xlink="https://www.w3.org/1999/xlink" version="1.1" id="Layer_1" x="0px" y="0px" viewBox="0 0 32 32" xml:space="preserve">
					<path xmlns="http://www.w3.org/2000/svg" d="M29.7,13.7H18.3V2.3C18.3,1,17.3,0,16,0c-1.3,0-2.3,1-2.3,2.3v11.4H2.3C1,13.7,0,14.7,0,16c0,1.3,1,2.3,2.3,2.3h11.4v11.4  c0,1.3,1,2.3,2.3,2.3c1.3,0,2.3-1,2.3-2.3V18.3h11.4c1.3,0,2.3-1,2.3-2.3S31,13.7,29.7,13.7z"></path>
				</svg>
            </span>    

      </button>
    </div>
  </div>

  <button class="btn--close slp-panzoom--icon slp-btn--close">
        <svg class="slp--close-icon"version="1.1" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" x="0px" y="0px"
            viewBox="0 0 512.001 512.001" style="enable-background:new 0 0 512.001 512.001; width:20px;height:20px" xml:space="preserve">
        <g>
            <g>
                <path d="M284.286,256.002L506.143,34.144c7.811-7.811,7.811-20.475,0-28.285c-7.811-7.81-20.475-7.811-28.285,0L256,227.717
                    L34.143,5.859c-7.811-7.811-20.475-7.811-28.285,0c-7.81,7.811-7.811,20.475,0,28.285l221.857,221.857L5.858,477.859
                    c-7.811,7.811-7.811,20.475,0,28.285c3.905,3.905,9.024,5.857,14.143,5.857c5.119,0,10.237-1.952,14.143-5.857L256,284.287
                    l221.857,221.857c3.905,3.905,9.024,5.857,14.143,5.857s10.237-1.952,14.143-5.857c7.811-7.811,7.811-20.475,0-28.285
                    L284.286,256.002z"/>
            </g>
        </g>
        </svg>
  </button>
<div class="js-slp-zoom-popup slp-zoom--popup ">
<div class="slp-zoom--tooltip">
 <p class="js-slp-zoom-description" ></p>
 <p class="js-slp-zoom-price slp-zoom---tooltip-price"></p>
  <div>
</div>
 
</section>



                </div>
        

<script type="text/javascript">
  var iSaleDescriptioncharacters = '500';
  var sReadMore = 'Read more';
  var sReadLess = 'Read less';
</script>

<div id="main_center_0_ctl02_dvSaleInformation" class="sale-information--container">
    
    <div class="sale-information-heading--wrapper js-sale-information-heading--wrapper col-xs-12 no-padding" data-toggle="collapse" data-target="#sale-overview" aria-expanded="true">
        
        <h5 class="heading-5 sale-information--title font_medium no-padding col-xs-10">Sale overview</h5>
         
        <div class="collapse--btn col-xs-2 no-padding">
            <span class="icon-theme--light pull-right plus--icon">
                <svg class="icon svg--plus svg--icons" xmlns="https://www.w3.org/2000/svg" xmlns:xlink="https://www.w3.org/1999/xlink" version="1.1" id="Layer_1" x="0px" y="0px" viewbox="0 0 32 32" xml:space="preserve">
					<path xmlns="http://www.w3.org/2000/svg" d="M29.7,13.7H18.3V2.3C18.3,1,17.3,0,16,0c-1.3,0-2.3,1-2.3,2.3v11.4H2.3C1,13.7,0,14.7,0,16c0,1.3,1,2.3,2.3,2.3h11.4v11.4  c0,1.3,1,2.3,2.3,2.3c1.3,0,2.3-1,2.3-2.3V18.3h11.4c1.3,0,2.3-1,2.3-2.3S31,13.7,29.7,13.7z"/>
				</svg>
            </span>
            <span class="icon-theme--light pull-right minus--icon">
                <svg class="icon svg--minus svg--icons" xmlns="https://www.w3.org/2000/svg" xmlns:xlink="https://www.w3.org/1999/xlink" version="1.1" id="Layer_1" x="0px" y="0px" viewbox="0 0 32 32" xml:space="preserve">
					<path xmlns="http://www.w3.org/2000/svg" d="M29.7,4.6H2.3C1,4.6,0,3.5,0,2.3S1,0,2.3,0h27.4C31,0,32,1,32,2.3S31,4.6,29.7,4.6z"/>
				</svg>
            </span>
        </div>
    </div>
    
    <div class="col-xs-12 no-padding sale-information--wrapper js-sale-information--wrapper collapse in" id="sale-overview" aria-expanded="true">
        <div class="col-xs-12 no-padding">
        
            <div id="main_center_0_ctl02_saledescription" class="sale-overview-image--container no-padding col-xs-12 col-sm-6 col-lg-7">
                <ul class="col-xs-12 no-padding">
                    
                    <li class="col-xs-12 sale-overview no-padding" id="sale-overview">
                        <div class="p--primary_large sale-overview--content" data-readmore="" aria-expanded="false">
                            <p>Christie's <em>Jewels Online July </em>auction features a selection of over 300 lots of fine jewelry ranging from Antique to Contemporary. Spanning all price points, this online sale highlights iconic designs by renown jewelry houses such as Bulgari, Cartier, David Morris, David Webb, Graff, JAR, Tiffany & Co., and Van Cleef & Arpels. This summer auction is sure to delight seasoned collectors and enthusiasts alike!</p>
                        </div>
                        <p class="read-more--wrapper" style="display:none">
                            <a href="#" class="read-more--btn p--primary_large">Read more</a>
                        </p>
                    </li>
                    
                    <li class="col-xs-12 sale-overview no-padding">
                        <p class="play-video--wrapper">
                            <a href="https://www.christies.com/features/Jewels-Online-10787-7.aspx?lid=1" class="btn btn--secondary play-video--btn">Explore Viewing Room</a>
                        </p>
                    </li>
                    
                </ul>
            </div>
        		
            
            
            <div class="sale-information no-padding col-xs-12 col-sm-5 col-lg-4" id="sale-information">
                <ul class="col-xs-12 no-padding">
                    <!-- If Auction Time is available -->
                    
                            
                            <li class="col-xs-12 col-md-10 no-padding">
                                <a class="add-to-calendar--link p--primary_large link underlined js-sale-information-add-to-calendar" role="button" data-toggle="modal" data-target="#basicModal" data-url="https://www.christies.com/AjaxPages/AddToCalendar.aspx?intsaleid=28905&lid=1" data-remote="https://www.christies.com/AjaxPages/AddToCalendar.aspx?intsaleid=28905&lid=1">
                                    Add to calendar
                                </a>
                            </li>
                            
                            <li class="col-xs-12 col-md-10 auction--wrapper no-padding">
                                <ul class="col-xs-12 no-padding">
                                    <li class="auction--title uppercase no-padding">
                                        Auction
                                    </li>
                                    <li class="auction--item no-padding">
                        
                            <hr style='margin: 1px 0 10px 0;'><span>Online</span>
                            
                            <span class='divider--left'>22 July - 6 August </span>
                            
                        
                            </li>
                                </ul>
                            </li>
                        

                    
                </ul>
            </div>
             
            
            <div class="no-padding col-xs-12 col-sm-12 col-lg-12 js-sale-overview-featured-content" style="display: none"></div>					
      
            <!-- Sale Image container - start -->
            		
            <!-- Sale Image container - end -->
            
            
            <div class="col-xs-12 contact-department no-padding" id="contact-department">
                <p class="contact-department--title uppercase">Contact the Specialist Department</p>
                <ul class="col-xs-12 no-padding">
            

            <li class='col-lg-4 col-md-4 col-sm-6 col-xs-12 contact-department--items'>
            <h2 class='p--primary_large font_medium contemporary-arts--title contact-department--subtitle'>Jewelry Department</h2>
            <div class="contact-department--item no-padding">
                <p class='p--primary_large font_medium specialist--title'>Caroline Ervin</p>
                
                <p class="p--primary_large specialist--email"><a onClick="commonFunctionality.activityMap(this, 'Email Specialist');" href='mailto:cervin@christies.com' title='Email Specialist'>cervin@christies.com</a></p>
                <a class="p--primary_large specialist--contact-number" onClick="commonFunctionality.activityMap(this, 'Call Specialist');" href='tel:+1 212 636 2307' title='Call Specialist' >Tel:+1 212 636 2307</a>
            </div>
            
            

            
            
            <div class="contact-department--item no-padding">
                <p class='p--primary_large font_medium specialist--title'>Edward Klopfer</p>
                
                <p class="p--primary_large specialist--email"><a onClick="commonFunctionality.activityMap(this, 'Email Specialist');" href='mailto:eklopfer@christies.com' title='Email Specialist'>eklopfer@christies.com</a></p>
                <a class="p--primary_large specialist--contact-number" onClick="commonFunctionality.activityMap(this, 'Call Specialist');" href='tel:+1 212 636 2318' title='Call Specialist' >Tel:+1 212 636 2318</a>
            </div>
            </li>
            
                </ul>
            </div>
            
            
        </div>
    </div>
</div>
        
        
            <div class="content-container-search">
            
                <div class="search-Result-Container result-Output">
                    <div class="result-Output-Right-pnl">
                        <div class="col-md-12 col-lg-12 col-xl-12 container lot-listings hidden-print" id="LotListings">
                            <div class="cc-loader--inner hidden"></div>
                        </div>
                    </div>
                </div>
            
                <div id="dvRelatedFeatures" style="display: none"></div>
                <div id="dvRelatedSales" style="display: none"></div> 
            
            </div>
        
        </li>
    </ul>
    
    
    <div id="dialog" style="display: none">
        
    </div>
    <div class="modal fade in" id="basicModal" tabindex="-1" role="dialog" aria-labelledby="basicModal" aria-hidden="true" style="display: none;">
        <div class="modal-dialog">
            <div class="modal-content load-modal-data"></div>
        </div>
    </div>
</div>


  </form>

  <footer>
    <div class='christies-footer'>
      <cc-global-footer></cc-global-footer>
    </div>

  </footer>

  <script src="/app_build/ccHeaderFooter.bundle.js"></script>

	  
<!-- ===========================================
OMNITURE CODE SECTION-B BEGIN
===========================================-->
<!-- SiteCatalyst code version: H.14.
Copyright 1997-2007 Omniture, Inc. More info available at
http://www.omniture.com -->
<script  type='text/javascript'>
/* You may give each page an identifying name, server, and channel on
the next lines. */
s.prop5 = 'New';
s.prop6 = 'Not Logged In/No Account';
s.eVar22 = 'English';

</script>
<!-- End SiteCatalyst code version: H.14. -->
<!-- ===========================================
OMNITURE CODE SECTION-B END
===========================================-->
<!-- Machine Name: MAC02 -->

		
	 
<!-- ===========================================
    This should be the last thing just before the END Body Tag
     OMNITURE CODE SECTION-C BEGIN
     ===========================================-->

<!--- DTM TAG -->
<script type="text/javascript" language="javascript">_satellite.pageBottom();</script>
<!--- DTM TAG -->
	 
<!-- ===========================================
    OMNITURE CODE SECTION-C END
     ===========================================-->
		
  
  
  <!--SLP-->
  <script type="text/javascript">
      var CSCInactivityTime = '900000';
      var rootSecurePath = 'https://www.christies.com/';
      var EnableCSCInactivityLogout = 'Y';
  </script>

  <script src="/Build/scripts/v3/slp-bundle.min.js?180720201156"></script>

    <script src="/js/BidContainerPopUp.js" type="text/javascript"></script>
    <script src="https://widgets.interests.christies.com/js/interestSaveDialog.js" id="interestSaveDialogJs" type="text/javascript"></script>
    <script type="text/javascript" src="/js/site.pages.timesbasedsalemessages.js"></script>
    <script src="/static/Scripts/SaleLanding/jSaleLanding.js"></script>
  

  <div id="dialog" style="display: none" class="container"></div>

  <script type="text/javascript">
      document.cookie = 'CurrentLanguage=; path=/; Domain=.www.christies.com; expires=' + new Date(0).toUTCString();
      document.cookie = 'CurrentLanguage=; path=/; expires=' + new Date(0).toUTCString();
  </script>

  
  <script>
      //TODO: Move API call into a js file
      //Call the API for headerfooter navigation
      var currentPage = '&CurrentPageURI=' +encodeURIComponent(window.location.href);
      var xhr = new XMLHttpRequest();
      xhr.open('GET', '/ChristiesAPIServices/Dotcomapis/api/GlobalNavigation/GetItems?LanguageCode=' + _val + currentPage + '&referrer=' + document.referrer, true );
      xhr.withCredentials = true;
      xhr.send();
      //When the api has returned its data, assign it to global variable and load the angular code
      xhr.onload = function (event) {
        window.ccHeaderFooterJson = JSON.parse(xhr.response);
        window.angular.bootstrap(document.getElementsByTagName('cc-global-header'), ['header']);
        window.angular.bootstrap(document.getElementsByTagName('cc-global-footer'), ['footer']);
      };

  </script>
  
  <link rel="stylesheet" type="text/css" href="/css/FastSignup.css?180720201156" />
  <script type="text/javascript" src="/js/LoadFastSignUp.js" defer></script>
  

</body>
</html>
`
var pattern = "{\n  total   `regex(\"\\\"totalItemsCount\\\": ([0-9]+),\")`\n  desc   `css(\"#sale-overview\");string()`\n  viewing `css(\"#sale-information > ul > li:nth-child(3) > ul > li.auction--item.no-padding\");string()`\n  address `css(\"#sale-information > ul > li:nth-child(2) > ul > li.auction--item.no-padding > p.p--primary_large.font_medium.auction-item--title\");string()`\noffline_time `css(\"#sale-information > ul > li:nth-child(2) > ul > li.auction--item.no-padding > p:nth-child(2)\")`\ninformation `css(\"#sale-information > ul\");string()`\n feature_content  `css(\"#dvFeatureContent > div > ul\");string()`\ncontact_us `css(\"#contact-department > ul > li:nth-child(1)\");string()`\nassociate_specialist `css(\"#contact-department > ul > li:nth-child(2)\");string()`\ncoordinator `css(\"#contact-department > ul > li:nth-child(2)\");string()`\n sale_number  `css(\"body > div.wrap-page > div.wrap-head > main > div.container > div > div.col-md-7.col-sm-6.col-xs-12.text-container > div.col-xs-12.nopadl.nopadr.body-copy-small\")`\nonline_begin_time `css(\"#auction-details-viewings > div.col-sm-12.col-md-12.auction-section.nopadlr > div.auction-start-time\")`\nonline_end_time `css(\"#auction-details-viewings > div.col-sm-12.col-md-12.auction-section.nopadlr > div.auction-end-time\")`\n}"

func TestParser(t *testing.T) {
	for {
		d, _ := Parser(content, pattern)
		if d.(map[string]interface{})["contact_us"] == nil || d.(map[string]interface{})["contact_us"].(string) == "" {
			fmt.Println("err")
			break
		}

	}

}
