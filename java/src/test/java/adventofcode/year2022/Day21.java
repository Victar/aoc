package adventofcode.year2022;

import adventofcode.BaseTest;
import lombok.Data;
import org.junit.Ignore;
import org.junit.Test;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.Map;

public class Day21 extends BaseTest {

	public static final int DAY = 21;

	@Ignore @Test public void runDownloadInput() throws Exception {
		downloadInput(DAY);
	}

	@Test public void runSilver() throws Exception {
		final ArrayList<String> data = readStringFromFile("year2022/day" + DAY + "/input.txt");
		Map<String, Monkey> monkeys = new HashMap<>();
		Monkey root = null;
		for (final String input : data) {
			Monkey current = new Monkey(input);
			monkeys.put(current.name, current);
		}
		System.out.println(monkeys.get("root").getNumber(monkeys, new HashMap<>()));
	}

	@Test public void runGold() throws Exception {
		final ArrayList<String> data = readStringFromFile("year2022/day" + DAY + "/input2.txt");
		Map<String, Monkey> monkeys = new HashMap<>();
		for (final String input : data) {
			Monkey current = new Monkey(input);
			monkeys.put(current.name, current);
		}
		boolean found = false;
		long start = 1L;
		long end = 2L;
		long offset = 0;
		while (!found) {
			long diffStart = getDiff(start + offset, monkeys);
			long diffEnd = getDiff(end + offset, monkeys);
			if (diffStart == 0 || diffEnd == 0) {
				System.out.println("Found following answers:");
				for (long j = offset - start; j <= offset + end; j++) {
					if (getDiff(j, monkeys) == 0) {
						System.out.println(j);
					}
				}
				found = true;
			}
			if (diffStart > 0 && diffEnd < 0 || diffStart < 0 && diffEnd > 0) {
				offset = offset + start;
				start = 1L;
				end = 2L;
			} else {
				start = start * 2;
				end = end * 2;
			}
		}
	}

	public long getDiff(long i, Map<String, Monkey> monkeys) {
		HashMap<String, Long> DP = new HashMap<>();
		String meName = "humn";
		Monkey root = monkeys.get("root");
		Monkey rootLeft = monkeys.get(root.monkeyLeft);
		Monkey rootRight = monkeys.get(root.monkeyRight);
		DP.put(meName, i);
		long leftNumber = rootLeft.getNumber(monkeys, DP);
		long rightNumber = rootRight.getNumber(monkeys, DP);
		//		System.out.println(i + " leftNumber:" + leftNumber + "  rightNumber:" + rightNumber + " diff " + (leftNumber - rightNumber));
		return leftNumber - rightNumber;
	}

	@Data class Monkey {

		String name;
		String operation;
		boolean isSimple;
		long simpleNumber;
		String monkeyLeft;
		String monkeyRight;
		String monkeyOperation;

		public Monkey(String input) {
			String[] arr = input.split(": ");
			this.name = arr[0];
			this.operation = arr[1];
			try {
				this.simpleNumber = Integer.parseInt(arr[1]);
				this.isSimple = true;
			} catch (Exception ignore) {
			}
			if (!isSimple) {
				String[] ops = this.operation.split(" ");
				monkeyLeft = ops[0];
				monkeyOperation = ops[1];
				monkeyRight = ops[2];
			}
		}

		public long getNumber(Map<String, Monkey> monkeys, Map<String, Long> DP) {
			if (DP.containsKey(this.name)) {
				return DP.get(this.name);
			}
			if (isSimple) {
				DP.put(this.name, this.simpleNumber);
				return simpleNumber;
			}
			long leftNumber = monkeys.get(monkeyLeft).getNumber(monkeys, DP);
			long rightNumber = monkeys.get(monkeyRight).getNumber(monkeys, DP);
			long result = 0;
			if ("+".equalsIgnoreCase(monkeyOperation)) {
				return leftNumber + rightNumber;
			}
			if ("-".equalsIgnoreCase(monkeyOperation)) {
				return leftNumber - rightNumber;
			}
			if ("*".equalsIgnoreCase(monkeyOperation)) {
				return leftNumber * rightNumber;
			}
			if ("/".equalsIgnoreCase(monkeyOperation)) {
				return leftNumber / rightNumber;
			}

			DP.put(this.name, result);
			return 0;
		}
	}

}
