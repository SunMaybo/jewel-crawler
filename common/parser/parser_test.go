package parser

import (
	"fmt"
	"testing"
)

var content=`

<!DOCTYPE html>
<html lang="zh-CN" dir="ltr">
  <head>
    <title>流感疫苗副作用大不能接种？腾讯新闻较真平台发布11月最全辟谣榜_腾讯新闻</title>
    <meta name="keywords" content="流感疫苗副作用大不能接种？腾讯新闻较真平台发布11月最全辟谣榜,流行性感冒,鳌拜,布洛芬,葛根粉,椰子油,谣言,辟谣,疫苗">
    <meta name="description" content="11月份腾讯新闻较真平台也依旧在全力终结谣言、播种知识。快来看看过去一个月中社会国际、食品安全、医疗健康领域的TOP3谣言，你信过哪个？想了解辟谣榜单全部内容？打开“腾讯较真辟谣”小程序，搜索“辟谣……">
    <meta name="apub:time" content="12/1/2020, 10:46:16 AM">
    <meta name="apub:from" content="default">
    <meta http-equiv="X-UA-Compatible" content="IE=Edge" />
<link rel="stylesheet" href="//mat1.gtimg.com/pingjs/ext2020/dcom-static/build/static/css/static.css" />
<!--[if lte IE 8]><meta http-equiv="refresh" content="0; url=/upgrade.htm"><![endif]-->
<!-- <meta name="sogou_site_verification" content="SYWy6ahy7s"/> -->
<meta name="baidu-site-verification" content="jJeIJ5X7pP" />
<link rel="shortcut icon" href="//mat1.gtimg.com/www/icon/favicon2.ico" />
<link rel="stylesheet" href="//vm.gtimg.cn/tencentvideo/txp/style/txp_desktop.css" />
<script src="//js.aq.qq.com/js/aq_common.js"></script>
<script>
    // 判断如果是动态底层不加载此JS逻辑 2020/1/19 -1
    if(location.href.indexOf('rain') === -1){
        (function(){
            var bp = document.createElement('script');
            var curProtocol = window.location.protocol.split(':')[0];
            if (curProtocol === 'https') {
                bp.src = 'https://zz.bdstatic.com/linksubmit/push.js';        
            }
            else {
                bp.src = 'http://push.zhanzhang.baidu.com/push.js';
            }
            var s = document.getElementsByTagName("script")[0];
            s.parentNode.insertBefore(bp, s);
        })();
    }
</script>
<script src="//mat1.gtimg.com/pingjs/ext2020/configF2017/5df6e3b3.js" charset="utf-8"></script>
<script src="//mat1.gtimg.com/pingjs/ext2020/configF2017/5a978a31.js" charset="utf-8"></script>
<script>window.conf_dcom = apub_5a978a31 </script><!--[if !IE]>|xGv00|61438c491c69d576aec9846de884f28b<![endif]-->
<!--[if !IE]>|xGv00|038d6e161753081e56c192d04873c65c<![endif]-->
    <script>window.DATA = {
		"article_id": "20201130A0G1JJ",
		"article_type": "0",
		"title": "流感疫苗副作用大不能接种？腾讯新闻较真平台发布11月最全辟谣榜",
		"abstract": null,
		"catalog1": "social",
		"catalog2": "social_shipinanquan",
		"media": "较真",
		"media_id": "5107513",
		"pubtime": "2020-11-30 21:48:28",
		"comment_id": "6243369951",
		"tags": "流行性感冒,鳌拜,布洛芬,葛根粉,椰子油,谣言,辟谣,疫苗",
		"political": 0,
		"artTemplate": null,
		"FztCompetition": null,
		"FCompetitionName": null,
		"is_deleted": 0,
		"cms_id": "20201130A0G1JJ00",
		"videoArr": []
}
      
    </script>
  </head>
  <body>
    <div id="TopNav"></div>
    <div id="topAd"></div>
    <div class="qq_conent clearfix">
      <div class="LEFT">
        <h1>流感疫苗副作用大不能接种？腾讯新闻较真平台发布11月最全辟谣榜</h1>
        <div class="content clearfix">
          <div class="LeftTool" id="LeftTool"></div>
          <!--内容-->
          <div class="content-article">
            <!--导语-->
            <p class="one-p">11月份腾讯新闻较真平台也依旧在全力终结谣言、播种知识。快来看看过去一个月中社会国际、食品安全、医疗健康领域的TOP3谣言，你信过哪个？</p>
            <p class="one-p">想了解辟谣榜单全部内容？打开“<strong>腾讯较真辟谣</strong>”小程序，搜索“<strong>辟谣榜</strong>”，较真妹等你哦~</p>
            <p class="one-p">11月社会国际类谣言TOP3</p>
            <p class="one-p"><strong>谣言一：鳌拜府邸应称其为“鳌府”</strong></p>
            <p class="one-p">热度：52.9K</p>
            <p class="one-p">正确解读：“鳌府”的叫法有误，原因是清朝爵位主要分为三个系统：宗室爵位、异姓功臣爵位以及蒙古爵位，鳌拜虽是一等超武公，但也只是在功臣系统中排名高，不是王爷或贝勒之类，其府邸就不能叫“府”。</p>
            <p class="one-p"><strong>谣言二：广澳高速连环车祸事故原因为一男子用千斤顶当警示牌</strong></p>
            <p class="one-p">热度：29.1K</p>
            <p class="one-p">正确解读：事故原因并非网传“用千斤顶当警示牌”所致，而是小货车碰撞路面桥梁伸缩缝后失控撞向前方正常行驶的蓝色小轿车。</p>
            <p class="one-p"><strong>谣言三：李鸿章挨一枪，大清少赔一亿两</strong></p>
            <p class="one-p">热度：13.5K</p>
            <p class="one-p">正确解读：《马关条约》的赔款数额最终由三亿两白银减为二亿两白银，其主要原因是日方当初提出的方案原本是作为会谈的基础提出来的，因此并非完全没有修正的余地，不是出于李鸿章挨了一枪的原因。</p>
            <p class="one-p"><img class="content-picture" src="//inews.gtimg.com/newsapp_bt/0/12844449145/1000">
            </p>
            <p class="one-p"> 11月食品安全类谣言TOP3</p>
            <p class="one-p"><strong>谣言一：椰子油是世界上最健康的油</strong></p>
            <p class="one-p">热度：63.9K</p>
            <p class="one-p">正确解读：椰子油是一种饱和脂肪酸含量很高的油脂，从目前国际主流意见推荐我们要限制饱和脂肪酸摄入的角度来看，椰子油不仅不是最健康的油脂，甚至可以说是一种不健康的油脂，还是要少吃。</p>
            <p class="one-p"><strong>谣言二：高蛋白饮食致癌</strong></p>
            <p class="one-p">热度：50.8K</p>
            <p class="one-p">正确解读：“高蛋白饮食致癌”没有任何明确科学依据。之所以会有该说法流传，是因为有人把一些在实验室精心设计发生的动物实验结果，错误的套用在了人类身上。比起饮食中蛋白质的含量高低，我们更应该关注蛋白质的来源。</p>
            <p class="one-p"><strong>谣言三：吃葛根粉可以补充雌激素</strong></p>
            <p class="one-p">热度：49.3K</p>
            <p class="one-p">正确解读：葛根中含有比较丰富的异黄酮类化合物，但它跟真正的雌激素性质和活性差距很大。我们吃的是葛根粉等相关食品而不是提取出来的葛根素，相应的功效更会“弱化”很多，对于雌激素正常的朋友来说无需担心吃葛根粉会给自己补充了“雌激素”。</p>
            <p class="one-p"><img class="content-picture" src="//inews.gtimg.com/newsapp_bt/0/12844451110/1000">
            </p>
            <p class="one-p"> 11月医疗健康类谣言TOP3</p>
            <p class="one-p"><strong>谣言一：流感疫苗副作用大会致死，不能接种</strong></p>
            <p class="one-p">热度：74.9K</p>
            <p class="one-p">正确解读：流感疫苗的安全性几乎是100%。所有的疫苗都存在不良反应，流感疫苗的不良反应包括注射部位出现酸痛发红、头痛、发烧、恶心等，但这些副作用都是暂时性的，也并不严重，不是什么大问题。没必要因副作用而拒绝流感疫苗。</p>
            <p class="one-p"><strong>谣言二：按摩能去除法令纹</strong></p>
            <p class="one-p">热度：60.0K</p>
            <p class="one-p">正确解读：法令纹的形成是皮肤松弛、皮下组织移位和骨质流失的共同结果。通过按摩“去法令纹”，无论是从下向上，还是从内向外，无论是提拉后固定5秒钟还是5分钟，都无法解决上述3方面问题，即便刚刚按摩后“看起来有变化”，也是暂时的，会很快恢复从前。</p>
            <p class="one-p"><strong>谣言三：镇痛药布洛芬副作用大不能用</strong></p>
            <p class="one-p">热度：11.3K</p>
            <p class="one-p">正确解读：布洛芬是非甾体抗炎药不良反应中相对较轻的一个，并且使用布洛芬的时长通常是退热不超过3天，镇痛不超过5天，这也就降低了长期使用布洛芬的风险性。总的来说布洛芬是很安全的，在正常用法用量下使用没有问题。</p>
            <p class="one-p"><img class="content-picture" src="//inews.gtimg.com/newsapp_bt/0/12844451917/1000">
            </p>
            <p class="one-p"><img class="content-picture" src="//inews.gtimg.com/newsapp_bt/0/12837514775/1000">
            </p>
            <div id="Status"></div>
          </div>
        </div>
        <div id="Comment"></div>
        <div id="Recomend"></div>
      </div>
      <div class="RIGHT" id="RIGHT"></div>
    </div>
    <div id="bottomAd"></div>
    <div class="qq_footer" id="Foot"></div>
    <div id="GoTop"></div>
    <script src="//mat1.gtimg.com/libs/jquery/1.12.0/jquery.min.js"></script>
<script type="text/javascript" src="//h5.ssp.qq.com/static/web/websites/pcnewsplugin/sspad_20200821.js" charset="utf-8"></script>
<script src="//mat1.gtimg.com/pingjs/ext2020/dc2017/dist/m_tips/tips.js" async></script>
<script src="//mat1.gtimg.com/pingjs/ext2020/dc2017/publicjs/m/ping.js" charset="gbk"></script>
<script src="//mat1.gtimg.com/pingjs/ext2020/2018/js/check-https-content.js"></script>
<script>
if(typeof(pgvMain) == 'function'){pgvMain();}
</script>
<script type="text/javascript" src="//imgcache.qq.com/qzone/biz/comm/js/qbs.js"></script>
<script type="text/javascript">
var TIME_BEFORE_LOAD_CRYSTAL = (new Date).getTime();
</script>
<script src="//ra.gtimg.com/web/crystal/v4.7Beta05Build050/crystal-min.js" id="l_qq_com" arguments="{'extension_js_src':'//ra.gtimg.com/web/crystal/v4.7Beta05Build050/crystal_ext-min.js', 'jsProfileOpen':'false', 'mo_page_ratio':'0.01', 'mo_ping_ratio':'0.01', 'mo_ping_script':'//ra.gtimg.com/sc/mo_ping-min.js'}"></script>
<script type="text/javascript">
if(typeof crystal === 'undefined' && Math.random() <= 1) {
  (function() {
    var TIME_AFTER_LOAD_CRYSTAL = (new Date).getTime();
    var img = new Image(1,1);
    img.src = "//dp3.qq.com/qqcom/?adb=1&dm=new&err=1002&blockjs="+(TIME_AFTER_LOAD_CRYSTAL-TIME_BEFORE_LOAD_CRYSTAL);
  })();
}
</script>
<style>.absolute{position:absolute;}</style>
<!--[if !IE]>|xGv00|bfa6be71716f6329ed6738978b6c1e2d<![endif]-->

<script>
var _mtac = {};
(function() {
    var mta = document.createElement("script");
    mta.src = "//pingjs.qq.com/h5/stats.js?v2.0.2";
    mta.setAttribute("name", "MTAH5");
    mta.setAttribute("sid", "500651042");
    var s = document.getElementsByTagName("script")[0];
    s.parentNode.insertBefore(mta, s);
})();
</script><!--[if !IE]>|xGv00|4164f7ee385140b403f6150542823185<![endif]-->
<script src="//mat1.gtimg.com/pingjs/ext2020/dcom-static/build/static/js/static.js"></script>
<!--[if !IE]>|xGv00|78fe8a44ba68d8b81e1f6f713a13b0c5<![endif]-->
  </body>
</html>
`
var pattern = "{\n  total   `regex(\"\\\"totalItemsCount\\\": ([0-9]+),\")`\n  desc   `css(\"#sale-overview\");string()`\n  viewing `css(\"#sale-information > ul > li:nth-child(3) > ul > li.auction--item.no-padding\");string()`\n  address `css(\"#sale-information > ul > li:nth-child(2) > ul > li.auction--item.no-padding > p.p--primary_large.font_medium.auction-item--title\");string()`\noffline_time `css(\"#sale-information > ul > li:nth-child(2) > ul > li.auction--item.no-padding > p:nth-child(2)\")`\ninformation `css(\"#sale-information > ul\");string()`\n feature_content  `css(\"#dvFeatureContent > div > ul\");string()`\ncontact_us `css(\"#contact-department > ul > li:nth-child(1)\");string()`\nassociate_specialist `css(\"#contact-department > ul > li:nth-child(2)\");string()`\ncoordinator `css(\"#contact-department > ul > li:nth-child(2)\");string()`\n sale_number  `css(\"body > div.wrap-page > div.wrap-head > main > div.container > div > div.col-md-7.col-sm-6.col-xs-12.text-container > div.col-xs-12.nopadl.nopadr.body-copy-small\")`\nonline_begin_time `css(\"#auction-details-viewings > div.col-sm-12.col-md-12.auction-section.nopadlr > div.auction-start-time\")`\nonline_end_time `css(\"#auction-details-viewings > div.col-sm-12.col-md-12.auction-section.nopadlr > div.auction-end-time\")`\n}"

func TestParser(t *testing.T) {
	fmt.Println(ParserArticleWithReadability(content,"https://www.qq.cn"))

}
