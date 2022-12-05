package adventofcode.year2022;

import adventofcode.BaseTest;
import org.junit.Test;

import java.util.ArrayList;
import java.util.Collections;
import java.util.List;

public class Day5 extends BaseTest {

	@Test public void runSilver() throws Exception {
		solve(true);
	}

	@Test public void runGold() throws Exception {
		solve(false);
	}

	public void solve(boolean reverse) throws Exception {
		final ArrayList<String> data = readStringFromFile("year2022/day5/input.txt");
		int size = data.get(0).length() / 4 + 1;
		List<List> all = new ArrayList<>();
		for (int i = 0; i < size; i++) {
			List list = new ArrayList();
			all.add(list);
		}
		boolean addToStack = true;
		for (String input : data) {

			if (input.startsWith(" 1 ")) {
				addToStack = false;
			}
			if (input.startsWith("move")) {
				doMove(input, all, reverse);
			}
			if (addToStack) {
				addToStack(input + " ", all);
			}
		}
		for (int i = 0; i < size; i++) {
			System.out.print(all.get(i).get(all.get(i).size() - 1));
		}
		System.out.println();
	}

	public void addToStack(String input, List<List> all) {
		for (int i = 0; i < all.size(); i++) {
			String cur = input.substring(i * 4, (i + 1) * 4);
			if (cur.contains("[")) {
				all.get(i).add(0, cur);
			}
		}
	}

	public void doMove(String input, List<List> all, final boolean reverse) {

		String[] arr = input.split(" ");
		int count = Integer.parseInt(arr[1]);
		List fromList = all.get(Integer.parseInt(arr[3]) - 1);
		List toList = all.get(Integer.parseInt(arr[5]) - 1);
		List take = fromList.subList(Math.max(fromList.size() - count, 0), fromList.size());

		if (reverse) {
			Collections.reverse(take);
		}

		int takeSize = take.size();
		List add = new ArrayList(take);

		for (int i = 0; i < takeSize; i++) {
			fromList.remove(fromList.size() - 1);

		}
		toList.addAll(add);
	}

}


