package adventofcode.year2022;

import adventofcode.BaseTest;
import com.google.gson.Gson;
import lombok.Data;
import org.junit.Ignore;
import org.junit.Test;

import java.util.ArrayList;
import java.util.Collections;
import java.util.List;

public class Day13 extends BaseTest {

	public static final int DAY = 13;

	@Ignore @Test public void runDownloadInput() throws Exception {
		downloadInput(DAY);
	}

	@Test public void runSilver() throws Exception {
		final ArrayList<String> data = readStringFromFile("year2022/day" + DAY + "/input.txt");
		int curPair = 0;
		int correctSum = 0;
		for (int i = 0; i < data.size(); i = i + 3) {
			curPair++;
			Signal left = new Signal(data.get(i));
			Signal right = new Signal(data.get(i + 1));
			if (left.compareTo(right) < 0) {
				correctSum += curPair;
			}
		}
		System.out.println(correctSum);
	}

	@Test public void runGold() throws Exception {
		final ArrayList<String> data = readStringFromFile("year2022/day" + DAY + "/input.txt");
		List<Signal> signalList = new ArrayList<>();
		for (int i = 0; i < data.size(); i = i + 3) {
			signalList.add(new Signal(data.get(i)));
			signalList.add(new Signal(data.get(i + 1)));
		}
		signalList.add(new Signal("[[2]]"));
		signalList.add(new Signal("[[6]]"));
		Collections.sort(signalList);
		int index2 = 0;
		int index6 = 0;
		for (int i = 0; i < signalList.size(); i++) {
			System.out.println(signalList.get(i));
			if (signalList.get(i).getId().equals("[[2]]")) {
				index2 = i + 1;
			}
			if (signalList.get(i).getId().equals("[[6]]")) {
				index6 = i + 1;
			}

		}
		System.out.println(index2 * index6);
	}

	@Data static class Signal implements Comparable {

		final ArrayList json;
		final String id;

		public Signal(String input) {
			this.id = input;
			this.json = new Gson().fromJson(input, ArrayList.class);
		}

		@Override public int compareTo(Object right) {
			if (right instanceof Signal) {
				return compare(this.json, ((Signal) right).json);
			}
			return 1;
		}

		int compare(Object left, Object right) {
			if (left instanceof Double && right instanceof Double) {
				return Double.compare((Double) left, (Double) right);
			}
			if (left instanceof ArrayList && right instanceof ArrayList) {
				final ArrayList arLeft = (ArrayList) left;
				final ArrayList arRight = (ArrayList) right;
				for (int i = 0; i < Math.min(arLeft.size(), arRight.size()); i++) {
					int compareItems = compare(arLeft.get(i), arRight.get(i));
					if (compareItems != 0) {
						return compareItems;
					}
				}
				return Integer.compare(arLeft.size(), arRight.size());
			}
			if (left instanceof ArrayList && right instanceof Double) {
				ArrayList<Object> rightAr = new ArrayList<>();
				rightAr.add(right);
				return compare(left, rightAr);
			}
			if (left instanceof Double && right instanceof ArrayList) {
				ArrayList<Object> leftAr = new ArrayList<>();
				leftAr.add(left);
				return compare(leftAr, right);
			}
			return 0;
		}
	}

}
