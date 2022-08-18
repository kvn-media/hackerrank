/* Weather Observation Station 4 */
SELECT COUNT(CITY) - COUNT(DISTINCT CITY) FROM STATION;

/* Weather Observation Station 5 */
SELECT DISTINCT CITY, LENGTH(CITY) FROM STATION ORDER BY LENGTH(CITY) ASC, CITY LIMIT 1;

SELECT DISTINCT CITY, LENGTH(CITY) FROM STATION ORDER BY -LENGTH(CITY) ASC, CITY LIMIT 1;

select city, length(city) from station
order by length(city),city asc
limit 1;
select city, length(city) from station
order by length(city) desc
limit 1;

/* Weather Observation Station 6 */
SELECT DISTINCT CITY FROM STATION
WHERE CITY REGEXP '^[aeiou]';
/* If our REGEXP is ^[...], that is just syntax to match all Strings that start with any character within the brackets, where the dots represent as many characters as we wish. Where we have the dot dot dot, I put the characters aeiou. The final regular expression matches all strrings that start with aeiou. 

^ is negation if it is inside brackets []. If it is the 1st character of a REGEXP it does something different (matches beginning of String). The same symbol can do 2 different things.
*/

/* Weather Observation Station 7 */
SELECT DISTINCT CITY FROM STATION
WHERE CITY REGEXP '[aeiou]$';
/* Func REGEXP 
the REGEXP let's you use strings to match patterns. Here, it matches all strings that end in a vowel.

Did this REGEXP is used only for matching vowel?
It can be used to match anything. But here, it's used to match vowels at the end of the string.

if they ask to print city name which start from the vowel so then which fuction we use?
You would then put a ^ at the front instead of a $ at the end
*/

/* Weather Observation Station 8 */
SELECT DISTINCT CITY FROM STATION
WHERE CITY REGEXP '^[aeiou].*[aeiou]$';

/* why can't we do like the below? CITY REGEXP '^[aeiou]$'
Nope, it's because the regex is wrong.
^  			// string's beginning
[aeiou]   		// a SINGLE vowel
$  			// end of string

Therefore this would match the vowels and only the vowels. Even a string like "aaa" would not match, because it's more than 1 char long.

The correct one is the one you used at the beginning because of the ".*" bit:
^			// start of string
[aeiou]			// a single vowel
.			// any characted...
*			// ...repeated any number of times
[aeiou]			// another vowel
$			// end of string

Why does not below condition work?

regexp_like (column_name, '^[aeiou]*[aeiou]$','i')
because you have . missing before . select distinct city from station where REGEXP_LIKE (lower(city),'^[aeiou].[aeiou]$');
*/

/* Weather Observation Station 9 */
SELECT DISTINCT CITY FROM STATION
WHERE CITY REGEXP '^[^aeiou]';
/*
Putting a ^ inside the closed bracket means something completely different than putting it outside the brackets. Putting it inside the brackets makes it match all characters EXCEPT the ones inside the bracket. So instead of writing [bcdfghjklmnpqrstvwxyz], we can write [^aeiou]

why did you include ^ outside the closed bracket
Putting a ^ inside the closed bracket means something completely different than putting it outside the brackets. Putting it inside the brackets makes it match all characters EXCEPT the ones inside the bracket. So instead of writing [bcdfghjklmnpqrstvwxyz], we can write [^aeiou]

yeah as i did write [^aeiou] this but the query failed .. it only worked after i included ^[^aeiou]. so this means query the names of all the city which are not starting from vowel s we could have just used [^aeiou] this .. but it didnt worked explain please
The first ^ is to make it match the beginning of the String. The second ^ is negation. You need both to work.

If you just did [^aeiou] it would be strings without aeiou. You need the extra ^ in ^[^aeiou] to make sure to match strings that dont BEGIN with aeiou.
*/

/* Weather Observation Station 10 */
SELECT DISTINCT CITY FROM STATION WHERE CITY REGEXP '[^aiueo]$';
/*
select distinct city from station where city REGEXP '[^AEIOU]$'; 
This should perfectly work guys, in MySQL

SELECT DISTINCT CITY FROM STATION WHERE CITY REGEXP '[^aiueo]$';
This one work perfectly. You should lower string of vowels.
regexp_like (city , '[aeiou]$ , 'i')) i is for insensetive case c is for case sensative
*/

/* Weather Observation Station 11 */
SELECT DISTINCT CITY FROM STATION
WHERE CITY REGEXP '^[^aeiou]|[^aeiou]$';
/*
HOW U USE EITHER OR IN THIS CODE
The symbol | is basically the "OR" you are referring to.

but what is the meaning of ^ b/w aeiou brackets EX -[^aeiou]
The ^ inside the brackets says that we can use any letter EXCEPT aeiou. If there was no ^ then it would be the 5 letters aeiou, which is not what we want.
*/

/* Weather Observation Station 12 */
SELECT DISTINCT(CITY)FROM STATION WHERE CITY REGEXP '^[^aeiou]' AND CITY REGEXP '[^aeiou]$';

/* Higher than 75 marks*/
select name from students where marks > 75 order by right(name,3), id ;
/*
purpose of id in the end?
Because: "If two or more students both have names ending in the same last three characters (i.e.: Bobby, Robby, etc.), secondary sort them by ascending ID."
*/

/* Employee names */
SELECT NAME FROM EMPLOYEE ORDER BY NAME;
select name from employee order by name asc;

/* Emplyee salaries */
SELECT name FROM Employee WHERE salary > 2000 AND months < 10 ORDER BY employee_id;

/* Weather observation station 16 */
SELECT ROUND(MIN(LAT_N), 4) FROM STATION WHERE LAT_N > 38.7780;

/* Weather observation station 17 */
select round(long_w, 4) from station group by long_w, lat_n having min(lat_n) > 38.7780 order by lat_n limit 1;

/* Weather observation station 18 */
SELECT ROUND(ABS(A-C)+ABS(B-D),4)
FROM
    (SELECT 
    MIN(LAT_N)  AS A,
    MIN(LONG_W) AS B,
    MAX(LAT_N)  AS C,
    MAX(LONG_W) AS D
    FROM STATION)
    AS T1;

/* Weather observation station 19 */    
SELECT
    ROUND(SQRT(
        POWER(MAX(LAT_N)  - MIN(LAT_N),  2)
      + POWER(MAX(LONG_W) - MIN(LONG_W), 2)
    ), 4)
FROM 
    STATION;

/* Type of triangle */
SELECT CASE
    WHEN 2 * GREATEST(A, B, C) >= (A + B + C) THEN "Not A Triangle"
    WHEN A = B AND A = C                      THEN "Equilateral"
    WHEN A = B OR A = C OR B = C              THEN "Isosceles"
                                              ELSE "Scalene"
    END
FROM TRIANGLES