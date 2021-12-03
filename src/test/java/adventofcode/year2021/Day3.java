package adventofcode.year2021;

import java.util.ArrayList;
import java.util.List;

import org.junit.Test;

import adventofcode.BaseTest;

public class Day3 extends BaseTest {

	@Test public void runSilver() throws Exception {
		final ArrayList<String> data = readStringFromFile("year2021/day3/input.txt");
		final List<List<Integer>> all = new ArrayList();
		int length = 0;
		for (final String input : data) {
			length = Math.max(length, input.length());
			final List<Integer> currentData = new ArrayList<>();
			for (int i = 0; i < input.length(); i++) {
				currentData.add(Integer.parseInt("" + input.charAt(i)));
			}
			all.add(currentData);
		}
		String result = "";
		String result1 = "";

		for (int i = 0; i < length; i++) {
			int count0 = 0;
			int count1 = 0;

			for (int j = 0; j < all.size(); j++) {
				final int currentNumber = all.get(j).get(i);
				if (currentNumber == 0) {
					count0++;
				} else {
					count1++;
				}
			}
			if (count0 > count1) {
				result = result + "0";
				result1 = result1 + "1";

			} else {
				result = result + "1";
				result1 = result1 + "0";
			}
		}
		System.out.println(Integer.parseInt(result, 2) * Integer.parseInt(result1, 2));

	}

	@Test public void runGold() throws Exception {
		final ArrayList<String> data = readStringFromFile("year2021/day3/input.txt");
		final List<List<Integer>> all = new ArrayList();
		int length = 0;
		for (final String input : data) {
			length = Math.max(length, input.length());
			final List<Integer> currentData = new ArrayList<>();
			for (int i = 0; i < input.length(); i++) {
				currentData.add(Integer.parseInt("" + input.charAt(i)));
			}
			all.add(currentData);
		}
		String result = "";
		String result1 = "";
		List<List<Integer>> filterList = all;

		for (int i = 0; i < length; i++) {
			filterList = filterList(filterList, i, true);
			if (filterList.size() == 1) {
				final StringBuilder b = new StringBuilder();
				filterList.get(0).forEach(b::append);
				result = b.toString();
			}

		}
		filterList = all;

		for (int i = 0; i < length; i++) {
			filterList = filterList(filterList, i, false);
			if (filterList.size() == 1) {
				final StringBuilder b = new StringBuilder();
				filterList.get(0).forEach(b::append);
				result1 = b.toString();
			}
		}
		System.out.println(Integer.parseInt(result, 2) * Integer.parseInt(result1, 2));

	}

	public List<List<Integer>> filterList(final List<List<Integer>> all, final int position, final boolean findMax) {
		int count0 = 0;
		int count1 = 0;
		for (int j = 0; j < all.size(); j++) {
			final int currentNumber = all.get(j).get(position);
			if (currentNumber == 0) {
				count0++;
			} else {
				count1++;
			}
		}
		final int maxCount = count0 > count1 ? 0 : 1;
		final List<List<Integer>> result = new ArrayList();

		for (int j = 0; j < all.size(); j++) {
			final int currentNumber = all.get(j).get(position);
			if (findMax && currentNumber == maxCount) {
				result.add(all.get(j));
			}
			if (!findMax && currentNumber != maxCount) {
				result.add(all.get(j));
			}
		}
		return result;

	}

}
