package adventofcode.year2020;

import java.math.BigInteger;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.LinkedHashMap;
import java.util.Map;

import org.apache.commons.lang3.StringUtils;
import org.junit.Ignore;
import org.junit.Test;

import adventofcode.BaseTest;

public class Day13 extends BaseTest {

	@Test public void singleCheck() {

	}

	@Test public void runSilver() throws Exception {
		final ArrayList<String> data = readStringFromFile("year2020/day13/input.txt");

		for (String input : data) {
			System.out.println(input);
		}
		Integer time = Integer.parseInt(data.get(0));
		String[] arrBus = StringUtils.split(data.get(1), ",x");
		final ArrayList<Integer> busNames = new ArrayList<Integer>();
		for (String input : arrBus) {
			busNames.add(Integer.parseInt(input));
		}
		boolean found = true;
		for (int i = time; found; i++) {
			for (Integer busName : busNames) {
				if (i % busName == 0) {
					System.out.println(busName);
					System.out.println(i);
					System.out.println(busName * (i - time));
					found = false;
				}
			}
		}
	}

	@Test
	@Ignore
	public void runGold() throws Exception {
		final ArrayList<String> data = readStringFromFile("year2020/day13/input_sample.txt");

		//		for (String input : data) {
		//			System.out.println(input);
		//		}
		String[] arrBus = StringUtils.split(data.get(1), ",");
		Map<BigInteger, BigInteger> busPostiton = new LinkedHashMap<>();
		int p = 0;
		for (int i = 0; i < arrBus.length; i++) {
			if (!"x".equals(arrBus[i])) {
				busPostiton.put(new BigInteger(arrBus[i]), BigInteger.valueOf(p));
			}
			p++;
		}

		/// Test Example (ni, ai)
		busPostiton = new LinkedHashMap<>();
		busPostiton.put(BigInteger.valueOf(3), BigInteger.valueOf(1));
		busPostiton.put(BigInteger.valueOf(4), BigInteger.valueOf(2));
		busPostiton.put(BigInteger.valueOf(5), BigInteger.valueOf(3));
		//		busPostiton.put(4,2);
		//		busPostiton.put(5,3);

		///

		BigInteger bigN = BigInteger.valueOf(1);
		for (Map.Entry<BigInteger, BigInteger> entry : busPostiton.entrySet()) {
			System.out.println(entry.getKey() + " => " + entry.getValue());
			bigN = bigN.multiply(entry.getKey());
		}
		System.out.println(bigN);

		//Mi Map
		Map<BigInteger, BigInteger> miMap = new LinkedHashMap<>();
		for (Map.Entry<BigInteger, BigInteger> entry : busPostiton.entrySet()) {
			miMap.put(entry.getKey(), bigN.divide(entry.getKey()));
		}
		System.out.println(miMap);


		Map<BigInteger, BigInteger> uiviMap = new LinkedHashMap<>();

		BigInteger sum =  BigInteger.valueOf(0);
		for (Map.Entry<BigInteger, BigInteger> entry : busPostiton.entrySet()) {
			BigInteger keyMi = miMap.get(entry.getKey());
			System.out.println(keyMi + "  " + entry.getKey());
			System.out.println("computeInverse: ( " + entry.getKey().intValue() + ": " +keyMi.intValue() +" )" + computeInverse(entry.getKey().intValue(), keyMi.intValue()));
		}
		System.out.println(uiviMap);
	}

	@Test
	public void runGold2() throws Exception {
		final ArrayList<String> data = readStringFromFile("year2020/day13/input.txt");

		String[] arrBus = StringUtils.split(data.get(1), ",");
		long l = 600691418730595l;
		System.out.println(l);
		int p = 0;
		Map<Long, Long> busPostiton = new LinkedHashMap<>();
		for (int i = 0; i < arrBus.length; i++) {
			if (!"x".equals(arrBus[i])) {
				Long x = Long.valueOf(arrBus[i]);
				long rem =  x-p;
				while(rem<0){
					rem+=x;
				}
				busPostiton.put(x,rem);
			}
			p++;
		}
		long[] rem = new long[busPostiton.size()];// { 1, 2, 3 };
		long[] num = new long[busPostiton.size()];// { 3, 4, 5 };


		int i=0;
		for (Map.Entry<Long, Long> entry : busPostiton.entrySet()) {
			System.out.println(entry);
			num[i] = entry.getKey();
			rem[i] = entry.getValue();
			i++;
		}

	}

	public long computeMinX(long[] rem, long[] num) {
		long product = 1;
		for (int i = 0; i < num.length; i++) {
			product *= num[i];
		}

		long[] partialProduct = new long[num.length];
		long[] inverse = new long[num.length];
		long sum = 0;

		for (int i = 0; i < num.length; i++) {
			partialProduct[i] = product / num[i];
			inverse[i] = computeInverse(partialProduct[i], num[i]);
			sum += partialProduct[i] * inverse[i] * rem[i];
		}
		return sum % product;
	}

	public long computeInverse(long a, long b) {
		long m = b, t, q;
		long x = 0, y = 1;
		if (b == 1) return 0;
		// Apply extended Euclid Algorithm
		while (a > 1) {
			// q is quotient
			q = a / b;
			t = b;
			// m is remainder now, process
			// same as euclid's algo
			b = a % b;
			a = t;
			t = x;
			x = y - q * x;
			y = t;
		}
		// Make x1 positive
		if (y < 0) y += m;
		return y;
	}

}

