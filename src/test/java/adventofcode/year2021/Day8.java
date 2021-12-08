package adventofcode.year2021;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.HashSet;
import java.util.Set;

import org.apache.commons.lang3.StringUtils;
import org.junit.Test;

import adventofcode.BaseTest;
import lombok.Data;

public class Day8 extends BaseTest {

	//abcdef
	public static final String S_0 = "1110111"; //6
	public static final String S_1 = "0010010"; //2
	public static final String S_2 = "1011101"; //5
	public static final String S_3 = "1011011"; //5
	public static final String S_4 = "0111010"; //4
	public static final String S_5 = "1101011"; //5
	public static final String S_6 = "1101111"; //6
	public static final String S_7 = "1010010"; //3
	public static final String S_8 = "1111111"; //7
	public static final String S_9 = "1111011"; //6

	public static final HashMap<String, Integer> S_N = createMap();

	private static HashMap<String, Integer> createMap() {
		final HashMap<String, Integer> result = new HashMap<>();
		result.put(S_0, 0);
		result.put(S_1, 1);
		result.put(S_2, 2);
		result.put(S_3, 3);
		result.put(S_4, 4);
		result.put(S_5, 5);
		result.put(S_6, 6);
		result.put(S_7, 7);
		result.put(S_8, 8);
		result.put(S_9, 9);

		return result;
	}

	@Test public void runSingle() throws Exception {
		System.out.println(Signal.decode("cdfeb", "acedgfb"));
	}

	@Test public void runGold() throws Exception {
		final ArrayList<String> data = readStringFromFile("year2021/day8/input.txt");
		int count = 0;
		final ArrayList<String> allMasks = new ArrayList(getMasks());
		for (final String input : data) {
			final String[] arr = StringUtils.split(input, "\\|");
			final String[] arr1 = StringUtils.split(arr[1], " ");
			final String[] arr2 = StringUtils.split(arr[0], " ");

			boolean found = false;
			for (int i = 0; i < allMasks.size() && !found; i++) {
				final int current = Signal.checkMask(arr1, arr2, allMasks.get(i));
				if (current > 0) {
					count += current;
					System.out.println("current: " + current);
					found = true;
				}
			}
		}
		System.out.println(count);
	}

	private Set<String> getMasks() {
		final char[] data = { 'a', 'b', 'c', 'd', 'e', 'f', 'g' };
		final int size = 7;
		final Set<String> masks = new HashSet<>();
		for (int i = 0; i < size; i++) {
			for (int j = 0; j < size; j++) {
				for (int k = 0; k < size; k++) {
					for (int l = 0; l < size; l++) {
						for (int m = 0; m < size; m++) {
							for (int t = 0; t < size; t++) {
								for (int n = 0; n < size; n++) {
									final String mask = "" + data[i] + data[j] + data[k] + data[l] + data[m] + data[t] + data[n];
									masks.add(mask);
								}
							}
						}
					}
				}
			}
		}
		return masks;
	}

	@Test public void runSilver() throws Exception {
		final ArrayList<String> data = readStringFromFile("year2021/day8/input.txt");
		int count = 0;
		for (final String input : data) {
			final String[] arr = StringUtils.split(input, "\\|");
			final String[] arr1 = StringUtils.split(arr[1], " ");
			for (final String signal : arr1) {
				final int size = signal.length();
				if (size == 2 || size == 4 || size == 3 || size == 7) {
					count++;
				}
			}
		}
		System.out.println(count);

	}

	@Data static class Signal {

		public static int checkMask(final String[] arr1, final String[] arr2, final String mask) {
			for (final String arr2s : arr2) {
				if (decode(arr2s, mask) < 0) {
					return -1;
				}
			}
			final int a = decode(arr1[0], mask);
			if (a < 0) {
				return -1;
			}
			final int b = decode(arr1[1], mask);
			if (b < 0) {
				return -1;
			}
			final int c = decode(arr1[2], mask);
			if (c < 0) {
				return -1;
			}
			final int d = decode(arr1[3], mask);
			if (d < 0) {
				return -1;
			}
			return a * 1000 + b * 100 + c * 10 + d;
		}

		static int decode(final String input, final String mask) {
			String result = "";
			int onecount = 0;
			for (int i = 0; i < mask.length(); i++) {
				if (input.contains("" + mask.charAt(i))) {
					result = result + "1";
					onecount++;
				} else {
					result = result + "0";
				}
			}
			if (input.length() != onecount) {
				return -1;
			}
			if (S_N.containsKey(result)) {
				return S_N.get(result);
			}
			return -1;
		}
	}

}
