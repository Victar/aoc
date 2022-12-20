package adventofcode.year2022;

import adventofcode.BaseTest;
import lombok.Data;
import org.junit.Ignore;
import org.junit.Test;

import java.util.ArrayList;

public class Day20 extends BaseTest {

	public static final int DAY = 20;

	@Ignore @Test public void runDownloadInput() throws Exception {
		downloadInput(DAY);
	}

	@Test public void runGold() throws Exception {
		runAny(811589153L, 10);
	}

	@Test public void runSilver() throws Exception {
		runAny(1L, 1);
	}

	public void runAny(long key, int repeat) throws Exception {
		final ArrayList<String> data = readStringFromFile("year2022/day" + DAY + "/input.txt");
		final ArrayList<Number> numbers = new ArrayList<>();
		for (int i = 0; i < data.size(); i++) {
			Number current = new Number(Long.parseLong(data.get(i)) * key, i);
			numbers.add(current);
		}
		for (int r = 0; r < repeat; r++) {
			for (int i = 0; i < numbers.size(); i++) {
				Number current = null;
				int removeIndex = -1;
				for (int j = 0; j < numbers.size(); j++) {
					if (numbers.get(j).getPosition() == i) {
						current = numbers.get(j);
						removeIndex = j;
						break;
					}
				}
				numbers.remove(removeIndex);
				int addIndex = (int) ((removeIndex + current.getValue()) % (numbers.size()) + numbers.size()) % (numbers.size());
				numbers.add(addIndex, current);
			}
		}
		int index = 0;
		for (int i = 0; i < numbers.size(); i++) {
			if (numbers.get(i).getValue() == 0) {
				index = i;
			}
		}
		long num1 = numbers.get((index + 1000) % numbers.size()).getValue();
		long num2 = numbers.get((index + 2000) % numbers.size()).getValue();
		long num3 = numbers.get((index + 3000) % numbers.size()).getValue();
		System.out.println("num1: " + num1 + " num2: " + num2 + " num3: " + num3);
		System.out.println(num1 + num2 + num3);

	}

	@Data static class Number {

		long value;
		int position;

		public Number(long value, int position) {
			this.value = value;
			this.position = position;
		}
	}

}
