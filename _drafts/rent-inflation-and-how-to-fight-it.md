---
layout: post
---
More and more people wish to live in metropolitan areas around the world. In pursuit of better jobs, more cultural opportunities and a better life in general they start to look for a living space but shortly after realise how expensive these places have become or they have to fear that they cannot afford they current home anymore in the near future. In this post I will take a look into the rental market economy, what rent controls exists and how they are performing to understand why we are in the situation we are in right now.

## Why do I think about in the first place?
For a while now I am interested in the economics of the rental market. Mainly because I moved to one of Germany's largest metropolitan areas, Hamburg. You can find more than 5 million people here, almost 2 million living in the city itself leading to a population density of **2430 citizens per square kilometer** (1). That of course leads to a problem people experiencing in many high-demand regions over the world like New York City, London or Berlin. Sky-rocketing rent prices. My wife and I for example started out in an newly built appartment which cost me well over 30% of my net income and was located 45min outside of the city center were I was working. Only through connections and a lot of luck were we able to find flat that is close to work and priced significantly under 30% net income. That was concerning to me, especially if you take into account that Hamburg, like many european cities, has rent control in place.

Looking at it I asked myself the questions *How does the rental market and especially rent control work?* and *How to fix the insane rent growth?* In the following post I take a look at the state of rent control research, which controls exist and what could be applied to my hometown.

## What is Rent Control?
The answer two this questions seems obvious at first. You restrict how rents can develop over time. But when taking a closer look you realise that it is a bit more complicated.
  - What is it is you restrict?
  - Is this rule applied in a specific area of city or the city as a whole?
  - Are all rental units restricted or are there exceptions?
What makes it even more complicated is the fact that there is a distinction between first generation rent controls created during the second world war and newer versions which fall under the category of second generation contols [2].

### First generation rent controls
The first generation can be sumed up at rent freeze. It was introduced during the war in the USA to prevent people taking advantage of families whos members were abroad. After the war ended, this restriction was mostly removed. One famous examle, New York City. Many areas removed it complity or like Europe enter into the second generation of rent controls. [2]

### Second generation rent controls
As stated above, the second generation is highly diversified. Under its umbrella fall regulations like:
  - rent increase per time unit
  - rent increase per new tenant
  - eviction laws
  - allowed rent differences in a specific area
And they are location variant meaning one city might decide upon a different set of rules than another. Or they have the same rules but with differences in they parameters (How much are you allowed to increase the rent per let's say year?) [2].

That, of course, makes it quite hard to measure how rent controls are performing and thus to quantify if they are beneficiary or not. Taking furthermore into account that his is a highly political and emotional topic as it directly affects peoples lifes [3] you can see that finding an object answer is a hard problem. Additional parameters influencing the markets are:
  - marco and local economics
  - local tax structure, especially for properties and land
  - real-estate cycle
  - ratio tenants to landlords which can influence the law process

## Is it any good?
As stated above, the answer is not clear. Depending on which research you read you might find that it works for a specific area or not. Looking add several effects of rent control you realise why that is the case [2], [4]:

When adding controls and restrictions landlords might transform their rent units into condos if the market price for an appartment is high enough. Thus they can offset their rental losses but at the same time the supply of rent units goes down. In this scenarios it makes also less economical sense for developers to build new rent units. If we now add exceptions to the controls the situation might get even worse as landlords will try to modify their property to be excempt from the restrictions. That could be the case if they for example upgrade the building to a higher quality of services. What happens than is a growth in rent equal or higher compared to a environment without rent control at least for in-migrants - people moving into the city - or citizens who need to change their appartment. Because, what also seems to happen when controls are put in place is a reduction in tenant mobility. That means, if you are living in a rent controlled appartment you are inlikely to move to a new place.

But this of course also has its benefits since people from disadvantaged social groups, e.g. low-income, are getting forced out of their homes by increasing rents. And of course, when leaving the rental development up to the market without regulations prices might surge in popular areas leaving it only with an homogenous population mostly made up from wealthy individuals.

A last point, which one should take into consideration is the "abandoning" of buildings. In this case, I am taalking about not renting out available units. A landlord could thus circumvent established rent controls and wait until he try to get an excemption or wait until a change in the politic sentiment leads to a reduction or complete removal of restrictions.

Taking the above information into account it is near to impossible to make general statements about the effectiveness of rent control. One rather has to look each case independently and even then you might miss or underestimate some parameters influencing the environment and skewing your findings.

## A case study and what to do against ever increasing rents?
To describe which solutions might work, how they work and what possible shortcomings they have we will take a look at Hamburg. It is the town which brought me to the point of writing this post in the first place. First, let's have a short look at the numbers and current restrictions. The cities rents increased between 22% to 39% between 2011 and 2018 depending on the size of the appartment [5]. At the same time the state-wide inflation was roughly 2% a year or ~15% in 7 years [6]. Taking into account that Hambuurg has rent controls in place means these were either not effect at all (rents are the same as without controls) or not effective enough (rents would be higher without controls). The following controls are in place [7], [8], [9]:
  - rents can only be increased to `average-area-rent + 10%`
  - landlords are not allowed to rent their flats on AirBnB without a certificate from the city. Furthermore, they are only allowed to rent it as holiday appartment for 8 weeks a year and every let has to be registered.
  - if buildings/rent units are vacant for too long landlords have to pay a fine of up to 50,000EUR. **Note**: I couldn't find an official document for this rule. I am also not sure about the timeframe you are alot to keep a rent unit vacant.

But these restriction have a couple of downsides. First, to make sure rent increases are under the declared limits tenants have report misbehaviour which produces a conflict. If you need/want a flat in Hamburg you will not report a landlord who asked for excess rents. Those who should report are not willing to do so. Another problem with that rule is its set of excemptions you can find in the resources. But not only the rent-increase restriction flawed. Vacant buildings need either be reported by citizens or get found by official. The problem here is that the bureau responsible for checking for vacant buildings is understaffed.

It seems like the only rule which seems to ork so far is the restriction on repurposing appartments as holiday domicils.But here the city pushed the responsibility to AirBnB.

### What are feaseble solutions?
Below we discuss already implemented and new solutions.

#### Penalizing empty office and rent units
A city could penalise landlords which own office or rent units which stay vacant longer than a limit `t_vacant`. This penalty could be a small fine (`x%` of the current rent) starting after a short time (3 months) and increases over time. To make sure the city actually knows when a unit is vacant a central databases is needed. There, landlords register their property and if a new tenant signs a contract she has to report that to the city anyway (Hamburg). Thus, only then the status would be set to *used*.

The problem here is we don't know automatically when a unit becomes vacant. Tenants might move within the same city and we could deduce that information but they might also leave the city. It would require an at least nation wide database. Another solution could be to require the landlord to report the vacancy but here she could just decide not to do it as is currently sometimes the case in Hamburg.

### Restrict Holiday appartments
As already discussed Hamburg enforces laws on rent unit usage. A tenant or landlord is only allowed to let his private flat as holiday appartment for 8 weeks a year and only for by registering every holidy rent. AirBnB (or any such platform) has to make sure that a landlord or tenant providing a flat has registered his flat and got a valid certificate [10]

When it comes to regulating the usage of rent units as holiday domicils effectiveness of reducing rent growth of course depends on the actual popularity of AirBnB and the like in the area. And one has to keep in mind that it adds new bureaycratic overhead to process all certificate and holiday-let requests.

### Restrict hotel density
Since tourism is growing with the popularity of cities so is the number of hotels accommodating the tourists. Depending on the city or area one might think about a upper limit for hotels per area unit or citizen number. But here it might be hard ti quantify what numbers make sense. Another approach could be to force hotel developers to also include appartments in their buildings making it a hotel/living space mix. This approach wouldn't create artificial limits and would add new flats to a city.

But the question is, would citizen would to live above or in a hotel building or in the areas hotels are created? Furthermore, the above ruling would immediately raise the question why it isn't applied to offices and other business units like shops.

### Ease construction of new buildings
Some metropolitan areas were able to keep their rent growth at inflation like levels by making it simple to construct new buildings [3]. Simple here means developers have a fixed set of rules they have to follow when building and as long as they obey nobody can object the construction. Also subsidizing land gives the right incentives for the creation of new buildings. Looking at our example of Hamburg we find that developers have a fixed set of rules but they also have to take the local community into account. They could resist an application for a new construction and represent a hard to estimate risk.

Of course we have to keep in mind that overruling local communities to reduce the risk of developers might not be politically and socially feasable.

### Higher density constructions
Continueing with the above point new constructions should also be incentivised to accommodate a higher number of people. This can be achieved by building closer - reducing the distance between houses and using more of the area - but might lead to sealed areas, meaning with no green [11]. Another proofen solution could be to build higher. But that means height restriction, which exist in Hamburg and other european cities, must be removed or at least eased.

A counter argument to this scenario would be the description of the often historic city image when building large buildings. One measurment to reduce that effect could be to force developers to keep the design of the lower levels in the same style of the sourrinding area. Another problem with higher density buildings is that they either targeted at low-income groups and lead to getto creation as can be seen in old "Plattenbau" villages. Or they target the opposite spectrum oof high-income people and lead again to the situation of a homogenous society. A desirable solution should enforce buildings with a heterogenous composition of people.

Tall buildings could also have a positive effect on the percentage of green in a city by incorporating new city gardening concepts [12].

### Subsidice the creation of new rent units
Staying with the large buildings idea one could ask the question how do you make sure the composition of people is heterogenues. Developers need to pay a high price to build large buildings like skyscrapers and have to offset their investments. One way could be to subsidize a percentage of the appartments when used for rent. There could be subsidiaries for property taxes or income taxes on rent. Furthermore, a push to modulare tower construction could help to reduce the cost [13].

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
