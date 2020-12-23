package parser

import (
	"fmt"
	"testing"
)

var content = `

<html lang="fr" dir="ltr">
   <head>
      <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
      <meta name="robots" content="index, follow">
      <meta http-equiv="x-ua-compatible" content="ie=edge">
      <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1">
      <title>This incredibly smartwatch challenges the biggest brands of connected watches!</title>
      <link rel="stylesheet" type="text/css" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/4.6.3/css/font-awesome.min.css">
      <link rel="stylesheet" type="text/css" href="https://stackpath.bootstrapcdn.com/bootstrap/3.3.7/css/bootstrap.min.css">
      <link rel="stylesheet" href="css/global.css" type="text/css">
      <link rel="shortcut icon" type="image/png" href="favicon.ico">
      
      <style>
	      #div-mobile {
		      display: block;
		      width: 100%;
		      background: #347e12;
		      height: 55px;
		      line-height: 35px;
		      position: fixed;
		      bottom: 5px;
		      font-size: 22px;
		      color: #fff;
		      text-align: center;
		      border-top: 5px solid #2d9c2d;
		      border-bottom: 5px solid #0f690f;
		      
	      }
	      @media screen and (min-width: 801px) {
			  #div-mobile {
			    display: none;
			  }
			  #div-desktop {
			    display: block;
			  }
			}
	      
      </style>
      
        </head>
   <body dir="ltr">
    	      <!--<div class="top-advertorial"><span class="top-advertorial-text">Advertorial</span></div>-->
      <div class="global-header">
         <div class="container">
            <p class="brand"></p>
            <p class="advertorial">Infomercial</p>
         </div>
      </div>
      <div id="header">
         <!--<div class="page-container">
            <p class="logo-text">The Tech Blog</p>
         </div>-->
      </div>
      <div class="page-container">
         <div class="page-row">
            <div class="page-content">
	            <h1 class="main-heading text-center" style="font-weight: 900;"><span>VIDEO:</span> <b style="font-weight: 900;">How a connected watch can assist you ! </b></h1>
	            <div class="img-small-center">
		               <div style="text-align: center;margin-top: -5px;margin-bottom: 10px;">
<img src="images/dar.png" style="
    width: 56px;
    transform: rotate(-90deg);
    ">
<img src="images/dar.png" style="
    width: 56px;
    transform: rotate(-90deg);
    "><img src="images/dar.png" style="
    width: 56px;
    transform: rotate(-90deg);
    "></div>
                       <div class="embed-container">
	                       
                     <iframe width="100%" height="315" src="https://www.youtube.com/embed/bDetitRv-A4?rel=0" frameborder="0" allow="accelerometer; autoplay; encrypted-media; gyroscope; picture-in-picture" allowfullscreen=""></iframe>
                       </div>
                       <p class="caption"><!--VIDEO DESCRIPTION HERE</p>-->
                       
                    
                   </div>
                   <div style="text-align: center;margin-top: -11px;margin-bottom: 10px;">
<img src="images/dar.png" style="
    width: 56px;
    transform: rotate(90deg);
    ">
<img src="images/dar.png" style="
    width: 56px;
    transform: rotate(90deg);
    "><img src="images/dar.png" style="
    width: 56px;
    transform: rotate(90deg);
    "></div>
                <div style="margin-top: 30px;">
               <h1 class="main-heading text-center">Is it really possible to get a high-performing connected watch at a low price? <i>"An exceptional manufacturing quality and a price which beats all competitors."</i></h1>
               <div class="col-xs-12 col-md-12 no-padding">
                  <div class="rating-stars-container">
                     <img src="images/rating-stars.png" class="rating-stars">
                  </div>
                  <div class="current-date">
                     02.01.2020 - 12:48 <span class="vertical-separator">|</span>
                  </div>
                  <div id="share-container">
                     <div class="box">
                        <a href="#" class="fb_btn" target="_blank">
                           <div class="text box1"><i class="fa fa-facebook"></i></div>
                        </a>
                     </div>
                     <div class="box">
                        <a href="#" class="tw_btn" target="_blank">
                           <div class="text box2"><i class="fa fa-twitter"></i></div>
                        </a>
                     </div>
                  </div>
               </div>
               <!--<div class="graybox">
                  <p>Je suis un Nomade Numérique âgé de 31 ans. Au cours des 7 dernières années sans domicile, j'ai visité plus de 50 pays. J'ai vécu pendant des mois dans des pays comme la Thaïlande, le Mexique, la Turquie, l'Espagne, le Nicaragua et l'Afrique du Sud. Si vous aimez voyager ou si vous souhaitez voyager plus souvent, vous aimerez ce gadget exclusif que je vais commenter.</p>
               </div> -->
               <div class="col-xs-12 col-md-12 no-padding">
                  <a id="headline-image" href="https://stalence-alawants.icu/click" rel="noreferrer" target="_blank"><img src="images/world-watch-old-people.jpg" class="img-responsive" style="margin:0;padding:0px"></a>
               </div>  
                    <h3>A new start-up has just launched a new generation connected watch. They have named it <a href="https://stalence-alawants.icu/click" target="_blank">eClever</a> !</h3> <p>The manufacturers of this revolutionary ‘smartwatch’ have made a real impression at the beginning of this year, with thousands of customers already convinced. 
The <a href="https://stalence-alawants.icu/click" target="_blank">eClever</a> is the first high-tech new generation connected watch which is  <strong>3 TIMES less expensive</strong> than its direct competitors. The pickiest customers have all fallen for its sleek design, the ergonomics of its body and its ease of use.
</p>
                
                    <p><strong>The sales have really skyrocketed all over the world…</strong></p><p> And we wanted to help you understand why and how this connected watch has become popular so quickly?!</p>
                    <p>It is a known fact that the big brands earn huge amounts profits from their sales. Every year, they persuade customers to purchase supposedly new connected watches with the same functionalities as the previous versions but always at a higher price… (almost the price of an iPhone).</p>
                    <p>We have asked ourselves: would we be able to manufacture a connected watch as performant as a Samsung Galaxy or other connected watches which are currently ‘dominating’ the technological market without breaking the bank ?</p>
                    <p>We have conducted our own research in order to find out why everybody started choosing the <a href="https://stalence-alawants.icu/click" target="_blank">eClever</a> over, for example, a Samsung connected watch.
                    <h2>What is it? </h2>
                    <p>It is called <strong>eClever</strong>and has already conquered the hearts of thousands of users who have swapped their watch (classical or connected) for this technological beauty… Not surprising, considering it’s very similar to other high-performing connected watches, except a lot cheaper!</p>
                    <img src="images/world-watch-2.png" class="img-responsive" style="margin:0;padding:0px">
                    <p> The <a href="https://stalence-alawants.icu/click" target="_blank">eClever</a> is distributed by a French start-up, which shares the same Chinese manufacturing factories as its direct competitors - the quality of the parts used in the fabrication process is similar to the quality of those used in a Samsung, Fitbit or any other connected watch currently dominant on the market. 
</p>
                    
                    <p>This savant mix of technology has been developed in the same manner as other more expensive models of connected watches. The responsiveness to the touch is truly amazing and this high-tech connected watch is extremely fast and fluid! The operating system used by this watch is completely optimised in order to prevent any bug.</p>
        <!--            
		 <div class="row">
			<div class="col col-lg-4 col-md-4 col-sm-4 col-xs-12 col-4">
				<img src="images/mom-translator-x.gif" class="img-fluid" style="margin-right: 10px;margin-bottom: 10px;width:100%;">
		    </div>
		    <div class="col col-lg-4 col-md-4 col-sm-4 col-xs-12 col-4">
			    <img src="images/mom-translator-x3.gif" class="img-fluid" width="95%" style="margin-right: 10px;margin-bottom: 10px;width:100%;">
			</div>
			<div class="col col-lg-4 col-md-4 col-sm-4 col-xs-12 col-4">
				<img src="images/mom-translator-x2.gif" class="img-fluid" width="95%" style="width:100%;margin-bottom: 10px;" >
		    </div>
		</div> -->
			
                    <h2>Why is the eClever so popular?</h2>
                    <div class="row">
			<div class="col col-lg-4 col-md-4 col-sm-4 col-xs-12 col-4">
				<img src="images/video-world-watch-1.gif" class="img-fluid" style="margin-right: 10px;margin-bottom: 10px;width:100%;">
		    </div>
		    <div class="col col-lg-4 col-md-4 col-sm-4 col-xs-12 col-4">
			    <img src="images/video-world-watch-2.gif" class="img-fluid" width="95%" style="margin-right: 10px;margin-bottom: 10px;width:100%;">
			</div>
			<div class="col col-lg-4 col-md-4 col-sm-4 col-xs-12 col-4">
				<img src="images/video-world-watch-3.gif" class="img-fluid" width="95%" style="width:100%;margin-bottom: 10px;" >
		    </div>
		</div>
                    <p>After a thorough observation of the company which manufactures the <a href="https://stalence-alawants.icu/click" target="_blank">eClever</a>, we have concluded that this start-up has focused on the most important things: </p>
                    <div style="padding-left: 20px;">
                    <p>✅<span>The first Fitness watch with a high-performing heart monitor.</span></p>
                    <p>✅<span>HD touchscreen with an exceptional lighting which adapts to the natural environmental light.</span></p>
                    <p>✅<span>A highly-reliable heart rate monitor. </span></p>
                    <p>✅<span>Sleep and fitness tracker, as well as heart monitor all integrated in a connected watch which is 100% waterproof. </span></p>
                    <p>✅<span>The eClever will help you track your progress and will encourage you to be more active.</span></p>
                    <p>✅<span>Personal vocal assistant integrated functionality which allows the user to answer calls, organise meetings and receive notifications (SMS, Facebook, WhatsApp, Gmail…).</span></p>
                    </div>
                    <img src="images/world-watch-1.jpg" class="img-responsive">
                    
                    <p>For most connected watches lovers, these are the most important characteristics. The connected watch has been designed in such a way as to ensure that you are always up to date and functionalities such as the heart monitor provides you with much-needed peace of mind. </p>
                    <p>The <a href="https://stalence-alawants.icu/click" target="_blank">eClever</a> also has an in-built phone assistance technology which allows everyone, especially the most vulnerable people, to access help quickly in case of an emergency. Amongst other advanced functionalities present in the eClever, are the localisation by GPS and wifi, reminders and alarms for medication as well as notifications for other tasks and important events. Last but not least, an application for the weather forecast! </p>
                    <img src="images/senior-ww-checking.jpg" class="img-responsive">
                    <h2>What are the other characteristics of the eClever?</h2>
                    <p>The eClever is full of surprises and useful basic options which are also present in more popular brand smartwatches : </p>
                    <div style="padding-left: 20px;">
                    <p>✅<span>Protocol Bluebooth 4.0, which ensures a permanent and stable connection which is essential for using various applications.</span></p>
                    <p>✅<span>Heart sensors designed to allow people who suffer from heart issues or anxiety to know when to slow down on physical activity</span></p>
                    <p>✅<span>Size: 44 * 38 * 10.7mm - Weight: 50g</span></p>
                    </div>
                    <img src="images/world-watch-carac.png" class="img-responsive">
                    <h2>What did our editorial team think of the watch? </h2>
                    <img src="images/world-watch-mickey.jpg" class="img-responsive">
                    <p><i>‘’An amazing manufacturing quality and a price which defies any rival watch. I can choose between various display menus and I personally prefer the Mickey one.’’ J</i></p>
                    <img src="images/world-watch-3.jpg" class="img-responsive">
                    <p><i>“What really stood out for me is the durability of the battery. I don’t have to charge it everyday compared to my old connected watch.”</i></p>
                    <img src="images/world-watch-4.jpg" class="img-responsive">
                    <p><i>“The first time I wore my watch to the pool, my friends thought I was crazy. That’s because they didn’t know that the eClever is COMPLETELY Waterproof :D”</i></p>
                    <h2>how much can eClever cost?</h2>
	       			<p>Our smartwatch specialist had the opportunity to pre-test the <a href="https://stalence-alawants.icu/click" target="_blank">eClever</a>, and according to prior knowledge and expertise, the price of the watch launched by the French start-up was between estimated to be between $300 and $400. Of course, that was based on no prior knowledge of the actual price.</p>
                    <p>He initially thought there was a website error when he learnt about the launching price of $ 69 !!!(including the current special offer)</strong></p>               
                    <p>This is definitely an extremely low price, considering the quality and performance of the smartwatch <a href="https://stalence-alawants.icu/click" target="_blank">eClever</a>. It ticks all the boxes and expected functionalities of well-known brands currently on the market! </p>
                   <h2>What do others think…</h2>
                
					                   <!-- Section: Testimonials v.3 -->
					<section class="team-section text-center my-5" style="margin-top: 15px;">
															
					  <!--Grid row-->
					  <div class="row text-center">
					
					    <!--Grid column-->
					    <div class="col-md-4 mb-md-0 mb-5">
					
					      <div class="testimonial">
					        <!--Avatar-->
					        <div class="avatar mx-auto">
					          <img src="images/antoine.jpg" class="rounded-circle z-depth-1 img-fluid">
					        </div>
					        <!--Content-->
					        <h4 class="font-weight-bold dark-grey-text mt-4">Antoine (48 ans)</h4>
					        <h6 class="font-weight-bold blue-text my-3">Department manager (supermarket)</h6>
					        <p class="font-weight-normal dark-grey-text">
					          <i>"I am salesman working at a renowned company and I think the eClever is the best gift I could have made to myself… It saves me so much time everyday…and to be honest, with my experience, I can tell that this watch is not expensive at all! I have ordered 3 of them last month, thank you so much!"</i></p>
					        <!--Review-->
					        <img src="images/5-stars.png" alt="stars" width="120px" style="float: center;margin-top: -1px;">
					      </div>
					
					    </div>
					    <!--Grid column-->
					
					    <!--Grid column-->
					    <div class="col-md-4 mb-md-0 mb-5">
					
					      <div class="testimonial">
					        <!--Avatar-->
					        <div class="avatar mx-auto">
					          <img src="images/claude.jpg" class="rounded-circle z-depth-1 img-fluid">
					        </div>
					        <!--Content-->
					        <h4 class="font-weight-bold dark-grey-text mt-4">Claude (67 years)</h4>
					        <h6 class="font-weight-bold blue-text my-3">retired</h6>
					        <p class="font-weight-normal dark-grey-text">
					          <i>"This watch makes my life so much easier on a daily basis. It is able to do a lot more than a traditional watch! While I am doing my exercises, I am immediately notified if it detects anything wrong with my body. So proud of my purchase!"</i></p>
					        <!--Review-->
					          <img src="images/5-stars.png" alt="stars" width="120px" style="float: center;margin-top: -1px;">
					      </div>
					
					    </div>
					    <!--Grid column-->
					
					    <!--Grid column-->
					    <div class="col-md-4">
					
					      <div class="testimonial">
					        <!--Avatar-->
					        <div class="avatar mx-auto">
					          <img src="images/lina.jpg" class="rounded-circle z-depth-1 img-fluid">
					        </div>
					        <!--Content-->
					        <h4 class="font-weight-bold dark-grey-text mt-4">Lina (36 years)</h4>
					        <h6 class="font-weight-bold blue-text my-3">Photographer</h6>
					        <p class="font-weight-normal dark-grey-text">
					          <i>"I wasn’t really sure if a connected watch would be useful and suitable for me… I didn’t even know if I would even use it until I decided to order it…and you know what? It’s safe to say that I can’t live without it and its various functionalities which make my life so easy!"</i></p>
					        <!--Review-->
					        <img src="images/5-stars.png" alt="stars" width="120px" style="float: center;margin-top: -1px;">
					      </div>
					
					    </div>
					    <!--Grid column-->
					
					  </div>
					  <!--Grid row-->
					
					</section>
					<!-- Section: Testimonials v.3 -->
                   <h2>How can it be so affordable? </h2>
                   <p>The start-up which has created this technological jewel has exclusively  focused on the elaboration of the product itself and not on the marketing and the opening of stores. Everything happens online! This has enabled them to slash the price of their product while at the same time allowing them to sell a high-performing connected watch at an extremely low price.</p>
                   <p>On the other side of the market spectrum are the big famous companies which spend millions in advertising. This has an direct and immediate impact on the products sold to the consumer.  </p>

                    <h2>Conclusion: should you purchase it or not? </h2>
                <p>The quality, performance and elegance of the <a href="https://stalence-alawants.icu/click" target="_blank">eClever</a>, are characteristics which render it exceptional. You will be a happy customer with a feeling of having bagged a bargain for an intelligent watch which costs between $300 and $500. This is an unbeatable price! The watch houses a battery with lasting durability; you will be able to use the watch for days and days on end. Our competitors cannot guarantee that. This offer is limited in time and while stock lasts; so do yourself a favour and invest in this intelligent beauty. You won’t regret it!</p>
               <p>To summarise: the <a href="https://stalence-alawants.icu/click" target="_blank">eClever</a>  does not pale in comparison to a Samsung, Fitbit or any other branded connected watch on the market. It is equally performant and aesthetic and a lot cheaper. </p>
               <h3>Here is a comparison of its most important characteristics compared to other connected watches currently on the market:</h3>
			   <img src="images/table-us.jpg" class="img-responsive">

            

			<h2>How can I order the eClever?</h2>

            <h3>Order the <strong>eClever</strong> smartwatch from its <a href="https://stalence-alawants.icu/click" rel="noreferrer" target="_blank">official website</a> .</h3>

                   
                         <p class="importantUpdate">For a limited time, this gadget is available with a <a href="https://stalence-alawants.icu/click" target="_blank">50% discount.</a></p>
                  

<div style="border:2px solid #059206;border-radius:5px;margin-bottom:25px">
						<p style="background:#059206;font-size:20px;font-weight:bold;padding:6px 12px;color:white">Advantages:</p>
						<ul>
							<li>Save huge amounts of money</li>
							<li>Stay connected, always</li>
							<li>Long-life battery.</li>

						</ul>
					</div>
					
<div style="border:2px solid #7b7b7b;border-radius:5px">
						<p style="background:#7b7b7b;font-size:20px;font-weight:bold;padding:6px 12px;color:white">Drawbacks:</p>
						<ul>
							<li>Available stock could run out soon</li>
						</ul> 
					</div>
			 <br>
			 
			 
                   
                   <a href="https://stalence-alawants.icu/click" rel="noreferrer" class="btnn btn-warning btn-full-width" target="_blank">Click here to claim a 50% off offer specially for our readers</a>
                   <p><center><em>The eClever is an amazing gift idea for you or for your loved ones!</em></center></p>
               </div>
                
                
                

               
               


            </div>
            <br style="clear:both">
            <div class="page-sidebar">
               <div class="sidebar1 sidebar-separators last">
                  <div class="title-side1"><b>Enjoy 50% off when you purchase your eClever:</b></div>
                  <a href="https://stalence-alawants.icu/click" rel="noreferrer" target="_blank">
                  <img src="images/world-watch-1.jpg" class="img-side-new">
                  </a>
                     <a class="btnn btn-warning btn-full-width" href="https://stalence-alawants.icu/click" target="_blank">
                     <div class="check-side">Check the product’s availability  >></div>
                  </a>
               </div>
               <div style="display:none;width:238px;height:229px;float:left"></div>
            </div>
       </div>      </div>
      <div id="footer">
         <div id="m_v" class="page-container">
            <p class="text-center">
               © 2020 <!--<a target="_blank" href="https://www.japantech.net/muama/en/003_kdc7/impressum.html" rel="noreferrer">Impressum</a>--> <a href="https://ecleverwatch.com/contact.html" target="_blank">CONTACT</a>
<a href="https://ecleverwatch.com/privacy.html" target="_blank">Privacy Policy</a>
 <a href="https://ecleverwatch.com/terms.php" target="_blank">Terms of Service</a>
<a href="https://ecleverwatch.com/contact.html" target="_blank">Contact</a>
            </p>
         </div>
      </div>
      <a href="https://stalence-alawants.icu/click" target="_blank"><div id="div-mobile">Check availability</div></a>
         <div class="disclaimer-footer">
             <br><br>
             <p>This is an advertisement and not a news article, a blog or a consumer protection update.<br><br>
                The illustrated story and the person represented on this page are fictional. It is based on the results obtained from people who have tested the product in real life. The results and comments published are general in nature. You might not get the same results when you use the product. This website may receive payment based on the clicks which led to the purchase of any product presented on the page. 
             </p>
             <br><br><!-- leave here for bottom bar -->
             <br><br><!-- leave here for bottom bar -->
             <br><br><!-- leave here for bottom bar -->
             <br><br><!-- leave here for bottom bar -->
         </div>
      <div id="bottombar">
         <center><a href="https://stalence-alawants.icu/click" class="btn btn-warning" target="_blank">vérifier la disponibilité &gt;&gt;&gt;</a></center>
      </div>
      <script src="https://code.jquery.com/jquery-2.2.4.min.js" integrity="sha256-BbhdlvQf/xTY9gja0Dq3HiwQF8LaCRTXxZKRutelT44=" crossorigin="anonymous"></script>
      <script src="js/jquery-scrolltofixed-min.js" type="text/javascript"></script>
      <script src="js/scripts.js" type="text/javascript"></script>
      <script type="text/javascript">
         $(document).scroll(function() {
         	$(window).scroll(function() {
         	var y = $(this).scrollTop();
         	var top_of_element = $("#warum").offset().top;
         	var bottom_of_element = $("#warum").offset().top + $("#warum").outerHeight();
         	var bottom_of_screen = $(window).scrollTop() + $(window).innerHeight();
         	var top_of_screen = $(window).scrollTop();
         	if (y > 3000) {
         		$("#bottombar").css({bottom:'0'});
         	} else {
         		$("#bottombar").css({bottom:'-90px'});
         	}
         	if ((bottom_of_screen > top_of_element) && (top_of_screen < bottom_of_element)){
         		$("#bottombar").css({bottom:'0'});
         	}
         	else if ((bottom_of_screen > top_of_element) && (top_of_screen > bottom_of_element)) {
         		$("#bottombar").css({bottom:'-90px'});
         	}
         	});
         });
        
        $(document).ready(function(){
            currentUrl = encodeURIComponent(location.href);
            $(".fb_btn").attr("href", "https://www.facebook.com/sharer/sharer.php?u=" + currentUrl);
            $(".tw_btn").attr("href", "https://twitter.com/share?url=" + currentUrl);
        });
      </script>
      <style>.bottom-fixed {
  bottom: 90px !important;
  top: auto !important;
}
.sidebar1 {
  float: left;
}
.bottombar {
  bottom: -90px;
}
</style>
   </body>
</html>

`
var pattern = "{\n  total   `regex(\"\\\"totalItemsCount\\\": ([0-9]+),\")`\n  desc   `css(\"#sale-overview\");string()`\n  viewing `css(\"#sale-information > ul > li:nth-child(3) > ul > li.auction--item.no-padding\");string()`\n  address `css(\"#sale-information > ul > li:nth-child(2) > ul > li.auction--item.no-padding > p.p--primary_large.font_medium.auction-item--title\");string()`\noffline_time `css(\"#sale-information > ul > li:nth-child(2) > ul > li.auction--item.no-padding > p:nth-child(2)\")`\ninformation `css(\"#sale-information > ul\");string()`\n feature_content  `css(\"#dvFeatureContent > div > ul\");string()`\ncontact_us `css(\"#contact-department > ul > li:nth-child(1)\");string()`\nassociate_specialist `css(\"#contact-department > ul > li:nth-child(2)\");string()`\ncoordinator `css(\"#contact-department > ul > li:nth-child(2)\");string()`\n sale_number  `css(\"body > div.wrap-page > div.wrap-head > main > div.container > div > div.col-md-7.col-sm-6.col-xs-12.text-container > div.col-xs-12.nopadl.nopadr.body-copy-small\")`\nonline_begin_time `css(\"#auction-details-viewings > div.col-sm-12.col-md-12.auction-section.nopadlr > div.auction-start-time\")`\nonline_end_time `css(\"#auction-details-viewings > div.col-sm-12.col-md-12.auction-section.nopadlr > div.auction-end-time\")`\n}"

func TestParser(t *testing.T) {
	fmt.Println(ParserArticleWithReadability(content, "https://www.news-gadget.com/smartwatch/en/4/dmn.html?cep=oSMFt8DGvo2h6nkVisN8e9hVabE36-3jVXkEHtG3SaoE54kB3BAt75ga_iIIGWXCt4cLH2OHKOm5UUnk9bnupe-nLJnUyIhZPDrywSdHwtV44vHlqDKOOd2tIHqbvBKy2n1fzR5atekXnMVxOleZvG7V8ABt0VV9v4Ht1okGjJlKyl3xUmJOgrZ4JTzFHkzxhwlNxSfRutEhMWf-vt4b4-7lLSsEb8QF4-9NcLNtT9801ggwFA-5K1KiUbT_MI7jnoT8Lv9dnlsXAnLsdtBToIuHr2y-jxYjO384Ve9T7cOxWcf4VN-ZuZFbszFtHZcf3wJO9H7vBmPpzwVnwS-A7W7xw-jcAfyWBEXOrLBl0eKIlu8WIcRHsU2M01lXJNEz9ap6rg2w4oZzhY_w572IXCx1M3fAQO2V4gA6OeC5O_qFIuczqMrwLhKkzRl3SVgg3wjKM4t7yXsV_qrmnNXFhE2cGxd7dS1xO-_qGCEqgOmiStDTNnZpUZJUoTEs2C5QRJC0hldeUlM73SInvPCsFE7lPlL5owZfv2ARs_0_zGjj9Hs-wjswi0kieVP73FQb_QypDHmpLMtdHVLkp9RgYYJkRmzJmPFCiXMvIQ-XWQLeFp1yUt_rgLDdEB0imRDRCa_7xRT_2gPTbUsbx-28SA&lptoken=16330765060971575219&ad_title=This%20Watch%20surprises%20the%20whole%20country.%20The%20price?%20Ridiculous!%20(only%20$&publisher_name=$publisher_name$&section_name=$section_name$&campaign_id=00d0d9b207304bc504a697f1d1c713f26d&ad_id=009f1add38baa5c9e9754243156bf3b37c&OutbrainClickId=$ob_click_id$&section_id=$section_id$&obOrigUrl=true"))
}
