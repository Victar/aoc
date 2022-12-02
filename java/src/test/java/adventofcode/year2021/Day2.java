package adventofcode.year2021;

import java.math.BigInteger;
import java.util.ArrayList;

import org.junit.Test;

import adventofcode.BaseTest;

public class Day2 extends BaseTest {

	@Test public void runSilver() throws Exception {
		final ArrayList<String> data = readStringFromFile("year2021/day2/input.txt");
		int x = 0;
		int y = 0;
		for (final String item : data) {
			final String[] arr1 = item.split(" ");
			final String direction = arr1[0];
			final int size = Integer.parseInt(arr1[1]);
			if ("forward".equals(direction)) {
				x += size;
			}
			if ("down".equals(direction)) {
				y += size;
			}
			if ("up".equals(direction)) {
				y -= size;
			}
		}
		System.out.println(x * y);
	}

	@Test public void runGold() throws Exception {
		final ArrayList<String> data = readStringFromFile("year2021/day2/input.txt");
		BigInteger x = BigInteger.valueOf(0);
		BigInteger y = BigInteger.valueOf(0);
		BigInteger aim = BigInteger.valueOf(0);

		for (final String item : data) {
			final String[] arr1 = item.split(" ");
			final String direction = arr1[0];
			final int size = Integer.parseInt(arr1[1]);
			if ("forward".equals(direction)) {
				x = x.add(BigInteger.valueOf(size));
				y = y.add(aim.multiply(BigInteger.valueOf(size)));
			}
			if ("down".equals(direction)) {
				aim = aim.add(BigInteger.valueOf(size));
			}
			if ("up".equals(direction)) {
				aim = aim.subtract(BigInteger.valueOf(size));
			}
		}
		System.out.println(x.multiply(y));
	}
}
