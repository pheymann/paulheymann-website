---
layout: post
---
More and more people wish to live in metropolitan areas around the world. In pursuit of better jobs, more cultural opportunities and a better life in general they start to look for a living space but shortly after realise how expensive these places have become or they have to fear that they cannot afford their current home anymore. In this post I will take a look into the rental market economy, what rent controls exists and how they effectiv they are to understand why we are in the situation we are in right now.

## Why do I think about this topic in the first place?
For a while now I am interested in the economics of the rental market. Mainly because I moved to one of Germany's largest metropolitan areas, Hamburg. You can find more than 5 million people here, almost 2 million living in the city itself leading to a population density of **2430 citizens per square kilometer** (1). That of course leads to a problem people experiencing in many high-demand regions over the world like New York City, London or Berlin. Sky-rocketing rent prices. My wife and I for example started out in an newly built appartment which cost us well over 30% of our net income and was located 45min outside of the city center were I was working. Only through connections and a lot of luck were we able to find a flat that is close to work and priced significantly under 30% net income. That was concerning to me, especially if you take into account that Hamburg, like many european cities, has rent control in place.

Looking at it I asked myself the questions *How does the rental market and especially rent control work?* and *How to behold the insane rent growth?* In the following post I take a look at the state of rent control research, which controls exist and what could be applied to my hometown.

## What is Rent Control?
The answer to this questions seems obvious at first. You restrict how rents can develop over time. But when taking a closer look you realise that it is a bit more complicated than that.
  - What is it is you restrict?
  - Is this rule applied in a specific area of city or the city as a whole?
  - Are all rental units restricted or are there exceptions?
What makes it even more complicated is the fact that there is a distinction between first generation rent controls created during the second world war and newer versions which fall under the category of second generation contols [2].

### First generation rent controls
The first generation can be sumed up as rent freeze. It was introduced in the USA during the war to prevent people taking advantage of families whos members were abroad. After the war ended, this restriction was mostly removed. Many areas removed it completely or like Europe enter into the second generation of rent controls. [2]

### Second generation rent controls
As stated above, the second generation is highly diversified. Under its umbrella fall regulations like:
  - rent increase per time unit
  - rent increase per new tenant
  - eviction laws
  - allowed rent differences in a specific area
And they are location variant meaning one city might decide upon a different set of rules than another. Or they have the same rules but with differences in their parameters (How much are you allowed to increase the rent per let's say year?) [2].

That, of course, makes it quite hard to measure how rent controls are performing and thus to quantify if they are beneficiary or not. Taking furthermore into account that it is a highly political and emotional topic - it directly affects peoples lifes - you can see that finding an object answer is a hard problem [3]. Additional parameters influencing the markets and therefore might skew results are:
  - marco and local economics: Do we have a recession (macro) or reported a large local company bankruptcy (GM in Detroit)?
  - local tax structure, especially for properties and land
  - real-estate cycle: Like every market real estate followes a cycle [14]
  - ratio tenants to landlords: Having more tenants might lead to pro-control politics and vis versa.

## Is it any good?
As stated above, the answer is not clear. Depending on which research you read you might find that it works for a specific area or not. Looking add several effects of rent control you realise why that is the case [2], [4]:

When adding controls and restrictions landlords might transform their rent units into condos if the market price for an appartment is high enough. Thus they can offset their rental losses but at the same time the supply of rent units goes down. In this scenarios it also makes less economical sense for developers to build new rent units. If we now add exceptions to the controls the situation might get even worse as landlords will try to modify their property to be exempt from restrictions. That could be the case if they for example upgrade the building to a higher quality of services. What happens than is a growth in rent equal or higher compared to a environment without rent control, at least for in-migrants - people moving into the city - or citizens who need to change their appartment. Because, what also seems to happen when controls are put in place is a reduction in tenant mobility. That means, if you are living in a rent controlled appartment you are inlikely to move to a new place. Thus, controls tend to favor residents.

But this, of course, also has its benefits since people from disadvantaged social groups, e.g. low-income, would be forced out of their homes by increasing rents. And when leaving the rental development up to the market without regulations prices might surge in popular areas leaving it only with an homogenous population mostly made up from wealthy individuals.

A last point, which one should take into consideration is vacancy of buildings and units. A landlord might try to circumvent established rent controls by waiting until an exemption kicks in or a change in the politic sentiment leads to a reduction or complete removal of restrictions.

Taking the above information into account it is close to impossible to make general statements about the effectiveness of rent control. One rather has to look at each case independently and even then you might miss or underestimate some parameters influencing the environment and skewing your findings.

## A case study and what to do against ever increasing rents
To describe which solutions might work, how they work and what possible shortcomings they have we will take a look at Hamburg. It is the town which brought me to write this post in the first place. First, let's have a short look at the numbers and current restrictions. The cities rents increased between 22% to 39% between 2011 and 2018 depending on the size of the appartment [5]. At the same time the state-wide inflation was roughly 2% a year or ~15% in 7 years [6]. Taking into account that Hamburg has rent controls in place means these were either not effect at all (rents are the same as without controls) or not effective enough (rents would be higher without controls). The following controls are in place [7], [8], [9], [15], [16]:
  - Rents can only be increased to `average-area-rent + 10%`
  - Landlords are not allowed to rent their flats on AirBnB without an allowens from the city. Furthermore, they are only allowed to rent it as holiday appartment for 8 weeks a year and every let has to be registered.
  - If buildings/rent units are vacant for more than 4 months landlords have to pay a fine.
  - Affordable rent units are build by the city to accomodate low-income families. To get one of these units you need to be under a certain income level to get voucher.
  - The city build rent units for all income groups.

But these restriction have a couple of downsides. First, to make sure rent increases are under the declared limits tenants have report misbehaviour which produces a conflict of interest. If you need/want a flat in Hamburg you will not report a landlord who asked for excess rents. Which of course reduces the propability of getting a flat. Those who should report are therefore not willing to do so. Another problem with that rule is its set of exemptions. But not only the rent-increase restriction is flawed. Vacant buildings need either be reported by citizens or needs to be found by officials. The problem here is that the bureau responsible for checking for vacant buildings is understaffed and most civilians do not care enough.

### What are feasible solutions?
Below we discuss different approaches. Some are already implemented in our case study and some are new.

#### Penalizing empty office and rent units
A city could penalise landlords which own office or rent units which stay vacant longer than a limit `t_vacant`. This penalty could be a small fine (`x%` of the current rent) starting after a short time (e.g. 3 months) and which increases over time. To make sure the city actually knows when a unit is vacant a central databases is needed. There, landlords register their property and if a new tenant signs a contract she has to report that to the city. The last part, tenant registration is already implemented in Hamburg. Thus, only if a tenant is living in a unit then the status would be set to *used*.

The problem here is we don't know automatically when a unit becomes vacant. Tenants might move within the same city and we could deduce that information but they might also leave the city. It would require an at least nation wide database to capure most movements - excluding emmigrants. Another solution could be to require the landlord to report the vacancy but here she could just decide not to do so as is currently the case in Hamburg.

### Restrict Holiday appartments
As already discussed Hamburg enforces laws on rent unit usage. A tenant or landlord is only allowed to let his private flat as holiday appartment for 8 weeks a year and only by registering every holidy rent. AirBnB (or any such platform) has to make sure that a landlord or tenant providing a flat has registered his flat and got a valid allowens [10].

Of course, the effectiveness of regulating the usage of rent units as holiday domicils to reduce rent growth depends on the actual popularity of AirBnB and the like in that area. And one has to keep in mind that it adds new bureaycratic overhead to process all allowens and holiday-let requests.

### Restrict hotel density
Since tourism is growing with the popularity of cities so is the number of hotels accommodating the tourists. Depending on the city or area one might think about a upper limit for hotels per area unit or citizen number. But here it might be hard to quantify what limits make sense and adding that limit adds artificial scarcity to the market what could lead to a price surge following the supply-demand curve. Another approach could be to force hotel developers to also include appartments in their buildings making it a hotel/living space mix. This approach wouldn't create artificial limits and would add new flats to the city.

But the question is, would citizen would like live above or in a hotel building or in the areas hotels are created? Furthermore, the above ruling would immediately raise the question why it isn't applied to offices and other business units like shops as well. Here, we would enter a fairness discussion which might skew the end-result to a point were it isn't beneficial anymore.

### Ease construction of new buildings
Some metropolitan areas were able to keep their rent growth at inflation like levels by making it simple to construct new buildings [3]. *Simple* here means developers have a fixed set of rules they have to follow when building and as long as they obey nobody can object the construction. Also subsidizing land gives the right incentives for the creation of new buildings. Looking at our example of Hamburg we find that developers have a fixed set of rules but they also have to take the local community into account. They could resist an application for a new construction and represent a hard to estimate risk.

Of course we have to keep in mind that overruling local communities to reduce the risk of developers might not be politically and socially feasible.

### Higher density constructions
Continueing with the above point new constructions should also be incentivised to accommodate a higher number of people. This can be achieved by building closer - reducing the distance between houses and using more of the area - but might lead to more sealed areas, meaning with no green [11]. Another proofen solution could be to build higher. But that means height restriction, which exist in Hamburg and other european cities, must be removed or at least eased.

A counter argument to this scenario would be the destruction or at least degregation of the often old, historic city image when building large, mordern buildings. One measurment to reduce that effect could be to force developers to keep the design of the lower levels in the same style of the sourrinding area. Another problem with higher density buildings is that they target specific social groups. In one case, they are planed for low-income families in the context of political decisions to provide affordable homes. But having a high density of these units might cause ghetto creation. On the other hand, speculators and developers might target the opposite spectrum of high-income people and lead again to a ghetto creation and therefore of homogenous societal groups within the city. A desirable solution should enforce buildings with a heterogenous composition of people. I have to see if I find the time to take a look into this topic and work out some resources. For now this paragraph is more an opinion than based on facts.

A final note on higher density buildings, these buildings could also have a positive effect on the percentage of green in a city by incorporating new city gardening concepts [12].

### Subsidice the creation of new rent units
Staying with the large buildings idea one could ask the question how do you make sure the composition of people is heterogenues. Developers need to pay a high price to build large buildings like skyscrapers and have to offset their investments. One way could be to subsidize a percentage of the appartments when used for low-income/median-income rent. There could be subsidiaries for property taxes or income taxes on rent. Of course, there is no guarantee that investors and developers plan for low/medium range units just because of subsidice. Adding enforcement by requiring minimum number of such units could help. 

Furthermore, a push to modulare tower construction could help to reduce the cost [13].

### Restrict real-estate speculation
One point I hadn't look into yet is real-estate speculation. Investors plan new buildings or buy and upgrade old ones not to accomodate the cities and residents needs but to improve return-on-investment. An examplary behaviour could be the destruction of low/medium price rent units to create high end appartments. To control speculation restrictions on land ownership could be put. Only if a person has citizenship she is allowed to by land. But land might be owned by companies and here it becomes tricky. I don't see a way how a city or country enforces a company structure that prevents foreign money to invest in the local market. Furthermore, it also runs under the assumption that local people and companies aren't speculating on real-estate. And last but not least, by restricting the inflow of investments a city might harm its own growth as new buildings might not get constructed.

### City led development
The city itself can enter the real-estate market and develop rent-units. By doing so it can provide living space for lower price rate than the actual market value. Hamburg is doing that for all income groups [15]. But we have to keep in mind that the city can lower rents to a point were the units are not profiatble anymore simple by subsidizing it and thus getting a market advantage. Private landlords might not be able to gain a profit in this scenario and either let their property decay or stop constructing new buildings. Which could lead to a reduction in supply and could drive prices up even more.

## How about no control
Taking an efficient market standpoint we could argue that no controls are needed and the market will, at least in the longrun, find its equillibrium at the fair price. But that raises a couple of remarks and questions:
  - Is that statement true under the condition that the real-estate market is plotically and socially influenced? As I stated at the beginning living space directly impact every person and his way of life. Irrational behaviour, therefore, might distort the market even over the long run.
  - Cities are complex social structures. Letting the ratio of different social groups get out of bounds because of surging rents might produce long lasting damages to the city and market itself.
  - Even if we assume that the market is effective in the longrun, inefficiencies most likely will create resentment, to say the least, when people getting forced out of their homes or are not able to find a living space. That might impact political decisions which will influence the market in return. And of course, it also begs the moral questions of letting groups of people suffer until the market is in equillibrium.

## Conclusion
Measuring the effectiveness of rent control is hard as the real-estate market depends on external influences which might skew the data and lead to wrong conclusions, and because it is a highly localised problem making generalisations is difficult. And therefore it is hard to say which approach might help in the short and/or long run. Looking at some of the growing metropolitan areas it seems that making construction easier and therefore providing more supply keeps rents at bay. And that is also complaince with how the market is supposed to worked.

## References
 - [1] [https://en.wikipedia.org/wiki/Hamburg](https://en.wikipedia.org/wiki/Hamburg)
 - [2] [https://pubs.aeaweb.org/doi/pdfplus/10.1257/jep.9.1.99](https://pubs.aeaweb.org/doi/pdfplus/10.1257/jep.9.1.99)
 - [3] [http://freakonomics.com/podcast/rent-control/](http://freakonomics.com/podcast/rent-control/)
 - [4] [https://ir.lib.uwo.ca/cgi/viewcontent.cgi?referer=&httpsredir=1&article=1001&context=economicsipsle_dp](https://ir.lib.uwo.ca/cgi/viewcontent.cgi?referer=&httpsredir=1&article=1001&context=economicsipsle_dp)
 - [5] [(GER) https://www.wohnungsboerse.net/mietspiegel-Hamburg/3195](https://www.wohnungsboerse.net/mietspiegel-Hamburg/3195)
 - [6] [(GER) https://de.statista.com/statistik/daten/studie/1046/umfrage/inflationsrate-veraenderung-des-verbraucherpreisindexes-zum-vorjahr/](https://de.statista.com/statistik/daten/studie/1046/umfrage/inflationsrate-veraenderung-des-verbraucherpreisindexes-zum-vorjahr/)
 - [7] [(GER) https://www.hamburg.de/mietenspiegel/4606594/mietpreisbremse/](https://www.hamburg.de/mietenspiegel/4606594/mietpreisbremse/)
 - [8] [(GER) https://www.hamburg.de/wohnraumschutz/11977066/wohnraumschutzgesetz-aenderungen-2019/](https://www.hamburg.de/wohnraumschutz/11977066/wohnraumschutzgesetz-aenderungen-2019/)
 - [9] [(GER) https://www.zeit.de/hamburg/2014-04/wohnung-leerstand-hamburg](https://www.zeit.de/hamburg/2014-04/wohnung-leerstand-hamburg)
 - [10] [https://www.airbnb.com/help/article/856/hamburg?locale=en](https://www.airbnb.com/help/article/856/hamburg?locale=en)
 - [11] [https://www.welt.de/debatte/kommentare/article152736886/Neue-Wohnungen-wir-muessen-wieder-bauen-wie-um-1900.html](https://www.welt.de/debatte/kommentare/article152736886/Neue-Wohnungen-wir-muessen-wieder-bauen-wie-um-1900.html)
 - [12] [https://www.researchgate.net/publication/321748735_Innovative_Urbanizing_City_Garden_and_Parks_in_the_Future_Skyscrapers](https://www.researchgate.net/publication/321748735_Innovative_Urbanizing_City_Garden_and_Parks_in_the_Future_Skyscrapers)
 - [13] [https://www.fullstackmodular.com/](https://www.fullstackmodular.com/)
 - [14] [https://onlinelibrary.wiley.com/doi/abs/10.1111/1540-6229.00772](https://onlinelibrary.wiley.com/doi/abs/10.1111/1540-6229.00772)
 - [15] [(GER) https://www.hamburg.de/neubau-in-hamburg/](https://www.hamburg.de/neubau-in-hamburg/)
 - [16] [(GER) http://www.landesrecht-hamburg.de/jportal/portal/page/bshaprod.psml?showdoccase=1&st=lr&doc.id=jlr-WoPflGHArahmen&doc.part=X&doc.origin=bs](http://www.landesrecht-hamburg.de/jportal/portal/page/bshaprod.psml?showdoccase=1&st=lr&doc.id=jlr-WoPflGHArahmen&doc.part=X&doc.origin=bs)
