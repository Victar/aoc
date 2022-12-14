package adventofcode.year2020;

import java.math.BigInteger;
import java.util.ArrayList;
import java.util.Collections;
import java.util.List;

import org.junit.Test;

import adventofcode.BaseTest;

public class Day9 extends BaseTest {

	@Test public void singleCheck() {

	}

	@Test public void runSilver() throws Exception {
		final ArrayList<BigInteger> data = readBigIntFromFile("year2020/day9/input.txt");
		final int preamble = 25;
		for (int input = preamble; input < data.size(); input++) {
			if (!isValidNumberBig(data.get(input), preamble, input, data)) {
				System.out.println(data.get(input));
			}
		}
	}

	//	21806024
	@Test public void runGold() throws Exception {
		final ArrayList<BigInteger> data = readBigIntFromFile("year2020/day9/input.txt");
		final BigInteger invalid = new BigInteger("21806024");

		for (int input = 0; input < data.size(); input++) {

			BigInteger currentSet = BigInteger.valueOf(0);
			for (int j = input; j < data.size(); j++) {
				currentSet = currentSet.add(data.get(j));
				if (currentSet.equals(invalid) && input != j) {
					System.out.println("found");
					final List<BigInteger> subArray = data.subList(input, j);
					System.out.println(subArray);
					Collections.sort(subArray);
					System.out.println(subArray);
					System.out.println(subArray.get(0).add(subArray.get(subArray.size() - 1)));
				}
				if (currentSet.compareTo(invalid) > 0) {
					//	System.out.println("more tann needed");
				}
			}
		}
	}

	public boolean isValidNumber(final int number, final int preamble, final int index, final ArrayList<Integer> data) {
		final List<Integer> subArray = data.subList(index - preamble, index);
		return isAnyTwoSumm(number, subArray);
	}

	public boolean isValidNumberBig(final BigInteger number, final int preamble, final int index, final ArrayList<BigInteger> data) {
		final List<BigInteger> subArray = data.subList(index - preamble, index);
		return isAnyTwoSummBig(number, subArray);
	}

	public boolean isAnyTwoSummBig(final BigInteger number, final List<BigInteger> data) {
		for (int i = 0; i < data.size(); i++) {
			for (int j = 0; j < data.size(); j++) {
				if (i != j && data.get(i).add(data.get(j)).equals(number)) {
					return true;
				}
			}
		}
		return false;
	}

	public boolean isAnyTwoSumm(final int number, final List<Integer> data) {
		for (int i = 0; i < data.size(); i++) {
			for (int j = 0; j < data.size(); j++) {
				if (i != j && data.get(i) + data.get(j) == number) {
					return true;
				}
			}
		}
		return false;
	}
}
